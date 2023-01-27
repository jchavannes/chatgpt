package api

import (
	"encoding/json"
	"fmt"
)

func FineTuneList(apiKey string) ([]FineTune, error) {
	resp, err := HttpRequest{
		Url:    UrlFineTunes,
		ApiKey: apiKey,
	}.Get()
	if err != nil {
		return nil, fmt.Errorf("%w; error fine tunes api request", err)
	}
	var respObj struct {
		Object string
		Data   []FineTune
	}
	if err := json.Unmarshal([]byte(resp), &respObj); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling fine tunes api response", err)
	}
	return respObj.Data, nil
}

type FineTuneCreateRequest struct {
	TrainingFile string `json:"training_file"`
}

func FineTuneCreate(apiKey, filename string) (*FineTune, error) {
	reqData, err := json.Marshal(FineTuneCreateRequest{
		TrainingFile: filename,
	})
	if err != nil {
		return nil, fmt.Errorf("%w; error json marshalling fine tune create request", err)
	}
	resp, err := HttpRequest{
		Url:    UrlFineTunes,
		ApiKey: apiKey,
		Data:   reqData,
	}.Post()
	if err != nil {
		return nil, fmt.Errorf("%w; error fine tune create api request", err)
	}
	var respFineTune = new(FineTune)
	if err := json.Unmarshal([]byte(resp), respFineTune); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling fine tunes create api response", err)
	}
	return respFineTune, nil
}
