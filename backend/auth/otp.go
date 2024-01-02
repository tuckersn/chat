package auth

import (
	"bytes"
	"image/png"

	"github.com/pquerna/otp/totp"
	"github.com/tuckersn/chatbackend/util"
)

func GenerateTOTP(username string) (string, []byte, error) {
	totpSecret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      util.Config.Auth.TokenIssuer,
		AccountName: username,
	})
	if err != nil {
		return "", make([]byte, 0), err
	}

	totpImage, err := totpSecret.Image(500, 500)
	if err != nil {
		return "", make([]byte, 0), err
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, totpImage)
	if err != nil {
		return "", make([]byte, 0), err
	}

	return totpSecret.Secret(), buf.Bytes(), nil
}
