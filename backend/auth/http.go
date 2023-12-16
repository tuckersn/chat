package auth

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/db"
)

func HttpAuth(c *gin.Context) (Token, error) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		return Token{}, errors.New("Authorization header required")
	} else if bearer[:7] != "Bearer " {
		return Token{}, errors.New("Authorization header must start with 'Bearer '")
	}

	tokenString := bearer[7:]
	token, err := ParseToken(tokenString)
	if err != nil {
		switch err.Error() {
		case "Token is expired":
		case "Key path not found":
			return Token{}, err
		default:
			log.Println("Unhandled error during auth (401): ", err)
			return Token{}, err
		}
	}

	return token, nil
}

func HttpAuthResponseHandled(c *gin.Context) (Token, error) {
	token, err := HttpAuth(c)

	if err == nil {
		return token, nil
	}

	switch err.Error() {
	case "Token is expired":
		c.JSON(401, gin.H{
			"error": "Token is expired",
		})
	case "Key path not found":
		c.JSON(400, gin.H{
			"error": "Malformed token, a key was not found in the token",
		})
	default:
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
	}
	return Token{}, err
}

func LoginUser(c *gin.Context, user db.RecordUser, metadata interface{}) (Token, error) {
	token, err := CreateToken(user.Username, user.Id, user.Admin)
	if err != nil {
		log.Println("Failed to create token", err)
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return Token{}, err
	}

	_, err = db.InsertLogin(user.Id, db.RECORD_LOGIN_ORIGIN_WEB, c.ClientIP(), metadata)
	if err != nil {
		log.Println("Failed to insert login", err)
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return Token{}, err
	}

	return token, nil
}
