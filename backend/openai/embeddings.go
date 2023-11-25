package openai

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/buger/jsonparser"
)

type Embedding struct {
	Index     int       `json:"index"`
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
}

type EmbeddingsResponse struct {
	Embeddings []Embedding `json:"data"`
	Vectors    []float64   `json:"raw_data"`
	Model      string      `json:"model"`
	Object     string      `json:"object"`
	Usage      struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

func GenerateEmbeddings(content string) (*EmbeddingsResponse, error) {
	var client = &http.Client{}
	var req, _ = http.NewRequest("GET", "https://api.openai.com/v1/embeddings", nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
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

	var embeddings EmbeddingsResponse
	err = json.Unmarshal(data, &embeddings)
	if err != nil {
		return nil, err
	}

	return &embeddings, nil
}
