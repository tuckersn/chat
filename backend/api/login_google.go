package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/db"
	"github.com/tuckersn/chatbackend/google"
	"github.com/tuckersn/chatbackend/util"
)

// HttpLoginGoogle godoc
// @Summary redirects browser to the Google OAuth consent screen
// @Description https://developers.google.com/identity/protocols/oauth2/web-server#httprest
// @Tags Login
// @Accept json
// @Produce json
// @Success 200
// @Router /login/google [get]
func HttpLoginGoogle(c *gin.Context) {
	csrf_token := util.RandomString(64, []string{util.ALPHABET_ALPHANUMERIC})
	err := db.InsertUserIdentityGoogleRequest(csrf_token)
	if err != nil {
		panic(err)
	}

	values := url.Values{}
	values.Add("client_id", google.GetGoogleAppId())
	values.Add("redirect_uri", "https://"+c.Request.Host+"/login/google/receive")
	values.Add("response_type", "code")
	values.Add("scope", strings.Join(google.OAUTH_SCOPES, " "))
	values.Add("access_type", "offline")
	values.Add("state", csrf_token)
	values.Add("include_granted_scopes", "true")
	values.Add("prompt", "select_account")

	c.Redirect(302, fmt.Sprintf("https://accounts.google.com/o/oauth2/v2/auth?%s", values.Encode()))
}

// HttpLoginGoogleReceive godoc
// @Summary Receives the response of the Google OAuth consent screen
// @Description https://developers.google.com/identity/protocols/oauth2/web-server#httprest
// @Tags Login
// @Accept json
// @Produce json
// @Success 200
// @Router /login/google/receive [get]
func HttpLoginGoogleReceive(c *gin.Context) {
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
		"client_id":     google.GetGoogleAppId(),
		"client_secret": google.GetGoogleAppSecret(),
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
		c.JSON(400, gin.H{
			"error": "Invalid code",
		})
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid code",
		})
		return
	}

	fmt.Println(string(respBody))

	var respJson map[string]interface{}
	err = json.Unmarshal(respBody, &respJson)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid code",
		})
		return
	}

	access_token, ok := respJson["access_token"].(string)
	if !ok {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid code",
		})
		return
	}

	refresh_token, ok := respJson["refresh_token"].(string)
	if !ok {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid code",
		})
		return
	}

	id_token, ok := respJson["id_token"].(string)
	if !ok {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid code",
		})
		return
	}

	fmt.Println(access_token)
	fmt.Println(refresh_token)
	fmt.Println(id_token)

	profile, err := google.GetProfile(access_token)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"error": "Invalid code",
		})
		return
	}

	fmt.Println(profile)

}
