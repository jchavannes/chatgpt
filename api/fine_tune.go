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
