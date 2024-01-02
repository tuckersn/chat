package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/auth"
	"github.com/tuckersn/chatbackend/db"
	"github.com/tuckersn/chatbackend/google"
	"github.com/tuckersn/chatbackend/util"

	gosql "database/sql"
)

// HttpLoginGoogleRedirect godoc
// @Summary HTTP redirect to Google's OpenID Connect (OAuth 2.0) consent screen
// @Description {art 1 of the HTTP redirect to Google's OpenID Connect (OAuth 2.0) consent screen
// https://developers.google.com/identity/protocols/oauth2/web-server#httprest
// @Tags Login
// @Accept json
// @Produce json
// @Success 200
// @Router /login/google [get]
func HttpLoginGoogleRedirect(c *gin.Context) {

	csrf_token := util.RandomString(64, []string{util.ALPHABET_ALPHANUMERIC})
	err := db.InsertUserIdentityGoogleRequest(csrf_token, c.Request.Host)
	if err != nil {
		panic(err)
	}

	values := url.Values{}
	values.Add("client_id", util.Config.Google.AuthAppId)
	values.Add("redirect_uri", "https://"+c.Request.Host+"/login/google/receive")
	values.Add("response_type", "code")
	values.Add("scope", strings.Join(google.OAUTH_ID_SCOPES, " "))
	values.Add("access_type", "offline")
	values.Add("state", csrf_token)
	values.Add("include_granted_scopes", "true")
	values.Add("prompt", "consent")

	c.Redirect(302, fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?%s", values.Encode()))
}

// HttpLoginGoogleRedirectReceive godoc
// @Summary Redirect URI receiving address for Googl'e OAuth 2.0 flow
// @Description Part 2 of the HTTP redirect to Google's OpenID Connect (OAuth 2.0) consent screen
// https://developers.google.com/identity/protocols/oauth2/web-server#httprest
// @Tags Login
// @Accept json
// @Produce json
// @Success 200
// @Router /login/google/receive [get]
func HttpLoginGoogleRedirectReceive(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")

	fmt.Println(code)
	fmt.Println(state)

	_, err := db.GetUserIdentityGoogleRequestDeleteAfter(state)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid csrf token",
		})
		return
	}

	reqBody := map[string]string{
		"code":          code,
		"client_id":     util.Config.Google.AuthAppId,
		"client_secret": util.Config.Google.AuthAppSecret,
		"redirect_uri":  "https://" + c.Request.Host + "/login/google/receive",
		"grant_type":    "authorization_code",
	}

	jsonData, err := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "https://oauth2.googleapis.com/token", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid code",
		})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": "Invalid google API request",
		})
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": "Invalid response body",
		})
		return
	}

	id_token_str, err := jsonparser.GetString(respBody, "id_token")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid id_token json",
		})
		return
	}

	id_token, err := google.VerifyGoogleIDToken(id_token_str)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid id_token when verifying",
		})
		return
	}

	if id_token.Iss != "https://accounts.google.com" {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid id_token, wrong issuer",
		})
		return
	}

	if id_token.Aud != util.Config.Google.AuthAppId {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid id_token",
		})
		return
	}

	if id_token.EmailVerified != true {
		log.Println(err)
		c.JSON(401, gin.H{
			"error": "Email not verified",
		})
		return
	}

	newUser := false
	var user db.RecordUser
	googleIdentity, err := db.GetUserIdentityGoogleByGoogleId(id_token.Sub)

	if err == gosql.ErrNoRows {
		// Create new user
		newUser = true
		username := util.RandomString(6, []string{util.ALPHABET_ALPHANUMERIC})

		totpSecret, totpImg, err := auth.GenerateTOTP(username)
		if err != nil {
			log.Println("Error generating TOTP", err)
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}

		encodedTotpImg := base64.RawStdEncoding.EncodeToString(totpImg)

		totpReqData, err := json.Marshal(db.RecordUserTotpRequestMetadata{
			Secret:      totpSecret,
			ImageData:   encodedTotpImg,
			RequestTime: time.Now().Unix(),
		})
		if err != nil {
			log.Println("Error marshalling TOTP request metadata", err)
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}

		metadataJson := []byte(`{}`)
		metadataJson, err = jsonparser.Set(
			metadataJson,
			totpReqData,
			db.METADATA_KEY_TOTP_REQUEST,
		)
		if err != nil {
			log.Println("Error setting TOTP request metadata", err)
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}

		//TODO: check if id_token.EmailVerified is true
		user, err = db.InsertUser(username, id_token.Name, id_token.Email, &totpSecret, metadataJson)
		if err != nil {
			log.Println("Error creating user indentity", err)
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}
		googleIdentity, err = db.InsertUserIdentityGoogle(user.Id, id_token.Sub)
		if err != nil {
			log.Println("Error creating Google identity", err)
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}

	} else if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return

	} else {
		// Get existing user
		user, err = db.GetUserById(googleIdentity.UserId)
		if err != nil {
			log.Println(err)
			c.JSON(500, gin.H{
				"error": "Internal server error",
			})
			return
		}
	}

	token, err := LoginUser(c, user, map[string]any{
		"google_id":  id_token.Sub,
		"user_agent": c.Request.UserAgent(),
		"url":        "https://" + c.Request.Host + "/login/google/receive",
	})
	if err != nil {
		return // Error already handled by LoginUser
	}

	// c.Redirect(302, fmt.Sprintf("https://%s/account/oauth/google?token=%s&newUser=%t", util.GetRedirectBaseUrl(), token.Signed, newUser))
	c.SetCookie("token", token.Signed, 60*60*24*365, "/", util.Config.Http.Host, false, true)
	c.Redirect(302, fmt.Sprintf("https://%s/account/oauth/google?token=%s&newUser=%t", util.GetRedirectBaseUrl(), token.Signed, newUser))
}

// HttpLoginGoogleDisconnect godoc
// @Summary [2FA] Removes your Google account information from your account.
// @Description [Two Factor Authentication Required] Removes your Google identity record from your account. This will prevent you from logging in with Google.
// @Tags Login
// @Accept json
// @Produce json
// @Success 200
// @Router /login/google/disconnect [delete]
func HttpLoginGoogleDisconnect(c *gin.Context) {
	user, err := auth.HttpAuthResponseHandled(c)
	if err != nil {
		return // error handled in auth.HttpAuthWithResponse
	}

	///TODO: consider moving to a multi-account system and using TOTP.

	err = db.DeleteIdentityGoogleByUserId(user.UserId)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}
