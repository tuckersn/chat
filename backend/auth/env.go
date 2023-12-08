package auth

import "os"

func GetTokenSecret() string {
	secret := os.Getenv("CR_TOKEN_SECRET")
	if secret == "" {
		panic("CR_TOKEN_SECRET not set")
	}
	return secret
}
