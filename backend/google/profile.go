package google

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ProfileResponseName struct {
	DisplayName string `json:"displayName"`
	FamilyName  string `json:"familyName"`
	GivenName   string `json:"givenName"`
}

type ProfileResponseEmailAddress struct {
	Value string `json:"value"`
}

type ProfileResponsePhoto struct {
	Url string `json:"url"`
}

type ProfileResponse struct {
	Names          *[]ProfileResponseName        `json:"names"`
	EmailAddresses []ProfileResponseEmailAddress `json:"emailAddresses"`
	Photos         *[]ProfileResponsePhoto       `json:"photos"`
}

func GetProfile(token string) (ProfileResponse, error) {
	req, err := PeopleAPIRequest(token, "GET", fmt.Sprintf(
		"/v1/people/me?personFields=%s",
		"names,emailAddresses,photos",
	), nil)
	if err != nil {
		return ProfileResponse{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ProfileResponse{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ProfileResponse{}, err
	}
	if resp.StatusCode != 200 {
		return ProfileResponse{}, errors.New("Unknown error, response code: " + resp.Status + " - " + string(respBody))
	}
	var profileResponse ProfileResponse
	err = json.Unmarshal(respBody, &profileResponse)
	if err != nil {
		return ProfileResponse{}, err
	}
	if len(profileResponse.EmailAddresses) == 0 {
		return ProfileResponse{}, errors.New("No email found")
	}
	return profileResponse, nil
}
