package gitlab

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/tuckersn/chatbackend/util"
)

func GetAccessToken(refreshToken string, redirect_uri string, verifier string) (string, error) {
	var client = &http.Client{}

	jsonData, err := json.Marshal(map[string]string{
		"redirect_uri":  redirect_uri,
		"client_id":     GetAppId(),
		"client_secret": GetAppSecret(),
		"code":          refreshToken,
		"grant_type":    "authorization_code",
		"code_verifier": verifier,
	})
	if err != nil {
		return "", err
	}

	var req, _ = http.NewRequest("POST", HostURL()+"/oauth/token", bytes.NewBuffer(jsonData))
	//req.Header.Set("Authorization", "Bearer "+GetAppSecret())
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	resp, err = client.Do(req)
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

func GetUserInfo(accessToken string) (string, error) {
	var client = &http.Client{}

	var req, _ = http.NewRequest("GET", HostURL()+"/api/v4/user", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	var resp *http.Response
	resp, err := client.Do(req)
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
