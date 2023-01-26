package api

import (
	"encoding/json"
	"fmt"
)

type CompletionRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature"`
	MaxTokens   int     `json:"max_tokens"`
}

func GetCompletion(apiKey string, prompt string) (*Completion, error) {
	req := CompletionRequest{
		Model:     ModelTextDavinci003,
		Prompt:    prompt,
		MaxTokens: 7,
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("%w; error json marshalling completion request", err)
	}
	resp, err := HttpRequest{
		Url:    UrlCompletions,
		Data:   reqData,
		ApiKey: apiKey,
	}.Post()
	if err != nil {
		return nil, fmt.Errorf("%w; error models api request", err)
	}
	var completion = new(Completion)
	if err := json.Unmarshal([]byte(resp), completion); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling models api response", err)
	}
	return completion, nil
}
