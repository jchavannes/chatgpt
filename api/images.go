package api

import (
	"encoding/json"
	"fmt"
)

func GetDefaultImageRequest(prompt string) ImageRequest {
	return ImageRequest{
		Prompt: prompt,
		N:      1,
		Size:   ImageSize1024,
	}
}

func ImageCreate(apiKey, prompt string) (*Image, error) {
	reqData, err := json.Marshal(GetDefaultImageRequest(prompt))
	if err != nil {
		return nil, fmt.Errorf("%w; error json marshalling image create request", err)
	}
	resp, err := HttpRequest{
		Url:    UrlImagesGen,
		Data:   reqData,
		ApiKey: apiKey,
	}.Post()
	if err != nil {
		return nil, fmt.Errorf("%w; error images create api request", err)
	}
	var respObj struct {
		Object string
		Data   []Image
		Error  *Error
	}
	if err := json.Unmarshal(resp, &respObj); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling images create api response", err)
	}
	if respObj.Error != nil {
		return nil, fmt.Errorf("error images create api response: %s - %s", respObj.Error.Type, respObj.Error.Message)
	}
	if len(respObj.Data) == 0 {
		return nil, fmt.Errorf("error no image returned")
	}
	return &respObj.Data[0], nil
}
