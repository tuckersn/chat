package google

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/golang-jwt/jwt/v5"
)

/* Google OAuth2 Constant Scope Values */
var OAUTH_ID_SCOPES = []string{
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/userinfo.profile",
	// "https://www.googleapis.com/auth/drive",
}

type jwk struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

var googlePublicKeys map[string]*rsa.PublicKey

func GooglePublicKeys() (map[string]*rsa.PublicKey, error) {
	if googlePublicKeys != nil {
		return googlePublicKeys, nil
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/certs")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	googlePublicKeys := make(map[string]*rsa.PublicKey)
	keysJson, _, _, err := jsonparser.Get(body, "keys")
	if err != nil {
		return nil, err
	}
	jsonparser.ArrayEach(keysJson, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		jwk := jwk{}
		err = json.Unmarshal(value, &jwk)
		if err != nil {
			return
		}
		googlePublicKeys[jwk.Kid], err = jwkToPublicKey(&jwk)
		if err != nil {
			return
		}
	})
	return googlePublicKeys, nil
}

func jwkToPublicKey(jwk *jwk) (*rsa.PublicKey, error) {
	if jwk.Kty != "RSA" {
		return nil, errors.New("invalid key type")
	}

	nb, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, err
	}
	eb, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, err
	}

	e := int(new(big.Int).SetBytes(eb).Int64())

	pub := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nb),
		E: e,
	}

	return pub, nil
}

type GoogleIdToken struct {
	JWT           jwt.Token
	Sub           string `json:"sub"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Iss           string `json:"iss"`
	Exp           int64  `json:"exp"`
	Iat           int64  `json:"iat"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
}

func VerifyGoogleIDToken(idTokenStr string) (GoogleIdToken, error) {
	keys, err := GooglePublicKeys()
	if err != nil {
		return GoogleIdToken{}, err
	}

	token, err := jwt.Parse(idTokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}

		key, ok := keys[kid]
		if !ok {
			return nil, errors.New("key not found")
		}

		return key, nil
	})

	if err != nil {
		return GoogleIdToken{}, err
	}

	if !token.Valid {
		return GoogleIdToken{}, errors.New("invalid token")
	}

	return GoogleIdToken{
		JWT:           *token,
		Sub:           token.Claims.(jwt.MapClaims)["sub"].(string),
		Azp:           token.Claims.(jwt.MapClaims)["azp"].(string),
		Aud:           token.Claims.(jwt.MapClaims)["aud"].(string),
		Iss:           token.Claims.(jwt.MapClaims)["iss"].(string),
		Exp:           int64(token.Claims.(jwt.MapClaims)["exp"].(float64)),
		Iat:           int64(token.Claims.(jwt.MapClaims)["iat"].(float64)),
		Name:          token.Claims.(jwt.MapClaims)["name"].(string),
		Picture:       token.Claims.(jwt.MapClaims)["picture"].(string),
		GivenName:     token.Claims.(jwt.MapClaims)["given_name"].(string),
		Locale:        token.Claims.(jwt.MapClaims)["locale"].(string),
		Email:         token.Claims.(jwt.MapClaims)["email"].(string),
		EmailVerified: token.Claims.(jwt.MapClaims)["email_verified"].(bool),
		AtHash:        token.Claims.(jwt.MapClaims)["at_hash"].(string),
	}, nil
}

func APIRequest(token string, method string, path string, body *io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, "https://www.googleapis.com"+path, *body)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return req, err
	}
	return req, nil
}

func PeopleAPIRequest(token string, method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, "https://people.googleapis.com"+path, body)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return req, err
	}
	return req, nil
}
