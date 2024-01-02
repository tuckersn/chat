package auth

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func HttpAuth(c *gin.Context) (Token, error) {
	tokenStr := c.GetHeader("Authorization")
	if len(tokenStr) > 7 {
		if tokenStr[:7] != "Bearer " {
			return Token{}, errors.New("Malformed Authorization header")
		}
		tokenStr = tokenStr[7:]
	} else {
		var err error
		tokenStr, err = c.Cookie("token")
		if err != nil {
			return Token{}, errors.New("Cookie 'token' or header 'Authorization' not found")
		}

	}

	token, err := ParseToken(tokenStr)
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
		log.Println("Unhandled error during auth (500): ", err)
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
	}
	return Token{}, err
}
