package api

import (
	"encoding/base64"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/auth"
	"github.com/tuckersn/chatbackend/db"
)

func LoginUser(c *gin.Context, user db.RecordUser, metadata interface{}) (auth.Token, error) {
	token, err := auth.CreateToken(user.Username, user.Id, user.Admin)
	if err != nil {
		log.Println("Failed to create token", err)
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return auth.Token{}, err
	}

	_, err = db.InsertLogin(user.Id, db.RECORD_LOGIN_ORIGIN_WEB, c.ClientIP(), metadata)
	if err != nil {
		log.Println("Failed to insert login", err)
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return auth.Token{}, err
	}

	return token, nil
}

// HttpLoginRecent godoc
// @Summary Gets your recent logins (up to 10)
// @Description gets a list of recent logins for provided token's associated user.
// @Tags Login
// @Accept json
// @Produce json
// @Success 200
// @Router /login/recent [get]
func HttpLoginRecent(r *gin.Context) {
	user, err := auth.HttpAuthResponseHandled(r)
	if err != nil {
		return // error handled in auth.HttpAuthWithResponse
	}

	logins, err := db.GetLoginsRecent(user.UserId, 10)
	if err != nil {
		panic(err)
	}

	r.JSON(200, logins)
}

// HttpTotpImageFile
// @Summary Gets your TOTP QR code image
// @Description gets a QR code image for your TOTP secret.
// @Tags Non-API-Routes
// @Router /asset/totp_request.png [get]
func HttpTotpImageFile(r *gin.Context) {
	user, err := auth.HttpAuthResponseHandled(r)
	if err != nil {
		return // error handled in auth.HttpAuthWithResponse
	}

	totpRequest, err := db.GetTotpRequestMetadata(user.Username)
	if err != nil {
		panic(err)
	}

	data, err := base64.RawStdEncoding.DecodeString(totpRequest.ImageData)
	if err != nil {
		panic(err)
	}

	r.Writer.Header().Set("Content-Type", "image/png")
	_, err = r.Writer.Write(data)
	if err != nil {
		r.AbortWithError(500, err)
		return
	}
}
