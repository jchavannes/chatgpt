package api

import (
	"encoding/json"
	"fmt"
	"sort"
)

func GetModelList(apiKey string) ([]Model, error) {
	resp, err := HttpRequest{
		Url:    UrlModels,
		ApiKey: apiKey,
	}.Get()
	if err != nil {
		return nil, fmt.Errorf("%w; error models api request", err)
	}
	var respObj struct {
		Object string
		Data   []Model
	}
	if err := json.Unmarshal(resp, &respObj); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling models api response", err)
	}
	sort.Slice(respObj.Data, func(i, j int) bool {
		return respObj.Data[i].Created > respObj.Data[j].Created
	})
	return respObj.Data, nil
}
