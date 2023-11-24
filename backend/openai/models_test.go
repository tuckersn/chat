package openai

import "testing"

func TestGetModels(t *testing.T) {
	models, err := GetModels()
	if err != nil {
		t.Error("ERROR", err)
	}
	if len(models) == 0 {
		t.Error("No models returned", models)
	}
}
