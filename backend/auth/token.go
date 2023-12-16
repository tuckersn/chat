package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/tuckersn/chatbackend/db"
	"github.com/tuckersn/chatbackend/util"
)

var RequiredFields = []string{"sub", "exp", "iat", "iss", "username", "display_name", "email", "admin"}

type Token struct {
	UserId      int32  `json:"sub"`
	Expires     int64  `json:"exp"`
	IssuedAt    int64  `json:"iat"`
	Issuer      string `json:"iss"`
	Audience    string `json:"aud"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	Admin       bool   `json:"admin"`
	JWT         jwt.Token
	Signed      string `json:"signed"`
}

func CreateToken(username string, user_id int32, admin bool) (Token, error) {
	user, err := db.GetUserById(user_id)
	if err != nil {
		return Token{}, err
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":          user_id,
		"exp":          time.Now().Add(time.Second * time.Duration(util.Config.Auth.TokenExpirySeconds)).Unix(),
		"iat":          time.Now().Unix(),
		"nbf":          time.Now().Unix(),
		"iss":          util.Config.Auth.TokenIssuer,
		"aud":          util.Config.Auth.TokenAudience,
		"username":     username,
		"display_name": user.DisplayName,
		"email":        user.Email,
		"admin":        admin,
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

func ParseToken(jwtStr string) (Token, error) {
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetTokenSecret()), nil
	})
	if err != nil {
		return Token{}, err
	}
	if !token.Valid {
		return Token{}, jwt.ErrSignatureInvalid
	}

	claims := token.Claims.(jwt.MapClaims)

	user_id, ok := claims["sub"].(float64)
	if !ok {
		return Token{}, jwt.ErrInvalidKey
	}

	expires, ok := claims["exp"].(float64)
	if !ok {
		return Token{}, jwt.ErrInvalidKey
	}

	issued_at, ok := claims["iat"].(float64)
	if !ok {
		return Token{}, jwt.ErrInvalidKey
	}

	issuer, ok := claims["iss"].(string)
	if !ok {
		return Token{}, jwt.ErrInvalidKey
	}

	audience, ok := claims["aud"].(string)
	if !ok {
		return Token{}, jwt.ErrInvalidKey
	}

	username, ok := claims["username"].(string)
	if !ok {
		return Token{}, jwt.ErrInvalidKey
	}

	email, ok := claims["email"].(string)
	if !ok {
		return Token{}, jwt.ErrInvalidKey
	}

	display_name, ok := claims["display_name"].(string)
	if !ok {
		return Token{}, jwt.ErrInvalidKey
	}

	admin, ok := claims["admin"].(bool)
	if !ok {
		return Token{}, jwt.ErrInvalidKey
	}

	return Token{
		UserId:      int32(user_id),
		Expires:     int64(expires),
		IssuedAt:    int64(issued_at),
		Issuer:      issuer,
		Audience:    audience,
		Username:    username,
		Email:       email,
		DisplayName: display_name,
		Admin:       admin,
		JWT:         *token,
		Signed:      jwtStr,
	}, nil
}
