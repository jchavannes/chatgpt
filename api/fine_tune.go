package api

import (
	"encoding/json"
	"fmt"
	"strings"
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
	if err := json.Unmarshal(resp, &respObj); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling fine tunes api response", err)
	}
	return respObj.Data, nil
}

func FineTuneCreate(apiKey, fileId string) (*FineTune, error) {
	reqData, err := json.Marshal(FineTuneCreateRequest{
		TrainingFile: fileId,
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
	if err := json.Unmarshal(resp, respFineTune); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling fine tunes create api response", err)
	}
	return respFineTune, nil
}

func FineTuneCancel(apiKey, fineTuneId string) (*FineTune, error) {
	if !strings.HasPrefix(fineTuneId, "ft-") {
		return nil, fmt.Errorf("invalid fine tune id")
	}
	resp, err := HttpRequest{
		Url:    UrlFineTunes + "/" + fineTuneId + "/cancel",
		ApiKey: apiKey,
	}.Post()
	if err != nil {
		return nil, fmt.Errorf("%w; error cancel fine tune api request", err)
	}
	var fineTune = new(FineTune)
	if err := json.Unmarshal(resp, fineTune); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling cancel fine tune api response", err)
	}
	return fineTune, nil
}

func FineTuneEvents(apiKey, fineTuneId string) ([]Event, error) {
	resp, err := HttpRequest{
		Url:    UrlFineTunes + "/" + fineTuneId + "/events",
		ApiKey: apiKey,
	}.Get()
	if err != nil {
		return nil, fmt.Errorf("%w; error fine tune events api request", err)
	}
	var respObj struct {
		Object string
		Data   []Event
	}
	if err := json.Unmarshal(resp, &respObj); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling fine tune events api response", err)
	}
	return respObj.Data, nil
}
