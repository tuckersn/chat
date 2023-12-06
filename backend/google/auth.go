package google

import (
	"io"
	"net/http"
)

/* Google OAuth2 Constant Scope Values */
var OAUTH_SCOPES = []string{
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/userinfo.profile",
	"https://www.googleapis.com/auth/drive",
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
