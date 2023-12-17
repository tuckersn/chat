package openai

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/buger/jsonparser"
)

type ModelResponse struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Creatable bool   `json:"creatable"`
	OwnedBy   string `json:"owned_by"`
}

func GetModels() ([]ModelResponse, error) {
	var client = &http.Client{}
	var req, _ = http.NewRequest("GET", "https://api.openai.com/v1/engines", nil)
	req.Header.Set("Authorization", "Bearer "+APIKey())
	var resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		return nil, errors.New("Unauthorized, check your OPENAI_API_KEY")
	} else if resp.StatusCode != 200 {
		return nil, errors.New("Unknown error, response code: " + resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	data, _, _, err = jsonparser.Get(data, "data")
	if err != nil {
		return nil, err
	}

	var models []ModelResponse
	err = json.Unmarshal(data, &models)
	if err != nil {
		return nil, err
	}

	return models, nil
}
