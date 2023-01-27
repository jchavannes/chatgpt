package api

import (
	"encoding/json"
	"fmt"
)

func GetDefaultCompletionRequest(prompt string) CompletionRequest {
	return CompletionRequest{
		Model:       ModelTextDavinci003,
		Prompt:      prompt,
		Temperature: 0,
		MaxTokens:   64,
	}
}

func GetCompletion(apiKey string, prompt string) (*Completion, error) {
	reqData, err := json.Marshal(GetDefaultCompletionRequest(prompt))
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
