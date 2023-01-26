package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	ModelTextAda001     = "text-ada-001"
	ModelTextDavinci003 = "text-davinci-003"
)

const (
	UrlCompletions = "https://api.openai.com/v1/completions"
	UrlFiles       = "https://api.openai.com/v1/files"
	UrlModels      = "https://api.openai.com/v1/models"
)

type HttpRequest struct {
	Url    string
	Data   []byte
	ApiKey string
}

func (r HttpRequest) Get() (string, error) {
	resp, err := r.do(http.MethodGet)
	if err != nil {
		return "", fmt.Errorf("%w; error get api request", err)
	}
	return resp, nil
}

func (r HttpRequest) Post() (string, error) {
	resp, err := r.do(http.MethodPost)
	if err != nil {
		return "", fmt.Errorf("%w; error post api request", err)
	}
	return resp, nil
}

func (r HttpRequest) do(method string) (string, error) {
	req, err := http.NewRequest(method, r.Url, bytes.NewReader(r.Data))
	if err != nil {
		return "", fmt.Errorf("%w; error creating api request", err)
	}
	req.Header.Set("Authorization", "Bearer "+r.ApiKey)
	if len(r.Data) > 0 {
		req.Header.Set("Content-Type", "application/json")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w; error executing api request", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%w; error reading api response", err)
	}
	return string(body), nil
}
