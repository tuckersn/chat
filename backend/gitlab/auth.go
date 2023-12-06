package gitlab

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/tuckersn/chatbackend/util"
)

func GetAccessToken(refreshToken string, redirect_uri string) (string, error) {

	var verifier, _ = GetCodeVerifier()
	var data = url.Values{}
	data.Set("client_id", GetAppId())
	data.Set("client_secret", GetAppSecret())
	data.Set("grant_type", "authorization_code")
	data.Set("code", refreshToken)
	data.Set("code_verifier", verifier)
	data.Set("redirect_uri", redirect_uri)

	var client = &http.Client{}
	urlStr := HostURL() + "/oauth/token?" + data.Encode()
	var req, _ = http.NewRequest("POST", urlStr, nil)
	//req.Header.Set("Authorization", "Bearer "+GetAppSecret())

	var resp, err = client.Do(req)
	if err != nil {
		return "", err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("Unknown error, response code: " + resp.Status + " - " + string(respBody))
	}

	return string(respBody), nil

}

func GetCodeVerifier() (string, string) {
	verifier := util.RandomString(128, []string{util.ALPHABET_URL_SAFE})
	h := sha256.New()
	h.Write([]byte(verifier))
	return verifier, base64.StdEncoding.EncodeToString(h.Sum(nil))
}
