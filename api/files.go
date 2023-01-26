package api

import (
	"encoding/json"
	"fmt"
	"sort"
)

func FileList(apiKey string) ([]File, error) {
	resp, err := HttpRequest{
		Url:    UrlFiles,
		ApiKey: apiKey,
	}.Get()
	if err != nil {
		return nil, fmt.Errorf("%w; error files api request", err)
	}
	var respObj struct {
		Object string
		Data   []File
	}
	if err := json.Unmarshal([]byte(resp), &respObj); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling files api response", err)
	}
	sort.Slice(respObj.Data, func(i, j int) bool {
		return respObj.Data[i].Filename < respObj.Data[j].Filename
	})
	return respObj.Data, nil
}
