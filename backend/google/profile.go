package google

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func GetProfile(token string) (string, error) {
	req, err := PeopleAPIRequest(token, "GET", fmt.Sprintf(
		"/v1/people/me?personFields=%s",
		"names,emailAddresses,photos",
	), nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
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
