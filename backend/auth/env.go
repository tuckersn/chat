package auth

import (
	"github.com/tuckersn/chatbackend/util"
)

func GetTokenSecret() string {
	secret := util.Config.Auth.TokenSecret
	if secret == "" {
		panic("CR_TOKEN_SECRET not set")
	}
	return secret
}
