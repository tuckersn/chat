package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tuckersn/chatbackend/db"
)

type Token struct {
	Username string `json:"username"`
	UserId   int32  `json:"user_id"`
	Admin    bool   `json:"admin"`
	JWT      jwt.Token
	Signed   string `json:"signed"`
}

func CreateToken(username string, user_id int32, admin bool) (Token, error) {

	user, err := db.GetUserById(user_id)
	if err != nil {
		return Token{}, err
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":     username,
		"user_id":      user_id,
		"display_name": user.DisplayName,
		"email":        user.Email,
		"admin":        admin,
		"exp":          time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	jwtString, err := jwt.SignedString([]byte(GetTokenSecret()))
	if err != nil {
		return Token{}, err
	}
	return Token{
		Username: username,
		UserId:   user_id,
		Admin:    admin,
		JWT:      *jwt,
		Signed:   jwtString,
	}, nil
}
