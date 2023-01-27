package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"sort"
	"strings"
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
	if err := json.Unmarshal(resp, &respObj); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling files api response", err)
	}
	sort.Slice(respObj.Data, func(i, j int) bool {
		return respObj.Data[i].Filename < respObj.Data[j].Filename
	})
	return respObj.Data, nil
}

func FileRetrieve(apiKey, fileId string) ([]byte, error) {
	if !strings.HasPrefix(fileId, "file-") {
		return nil, fmt.Errorf("invalid file id")
	}
	resp, err := HttpRequest{
		Url:    UrlFiles + "/" + fileId + "/content",
		ApiKey: apiKey,
	}.Get()
	if err != nil {
		return nil, fmt.Errorf("%w; error delete file api request", err)
	}
	return resp, nil
}

func FileUpload(apiKey, filename string) (*File, error) {
	fh, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%w; error opening file", err)
	}
	defer fh.Close()
	var bodyBuf = new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodyWriter.WriteField("purpose", "fine-tune")
	parts := strings.Split(filename, "/")
	fileWriter, err := bodyWriter.CreateFormFile("file", parts[len(parts)-1])
	if err != nil {
		return nil, fmt.Errorf("%w; error creating form file", err)
	}
	if _, err = io.Copy(fileWriter, fh); err != nil {
		return nil, fmt.Errorf("%w; error copying file to form", err)
	}
	bodyWriter.Close()
	resp, err := HttpRequest{
		Url:         UrlFiles,
		ApiKey:      apiKey,
		Data:        bodyBuf.Bytes(),
		ContentType: bodyWriter.FormDataContentType(),
	}.Post()
	if err != nil {
		return nil, fmt.Errorf("%w; error files upload api request", err)
	}
	var respFile = new(File)
	if err := json.Unmarshal(resp, respFile); err != nil {
		return nil, fmt.Errorf("%w; error json unmarshalling file upload api response", err)
	}
	return respFile, nil
}

func FileDelete(apiKey, fileId string) error {
	if !strings.HasPrefix(fileId, "file-") {
		return fmt.Errorf("invalid file id")
	}
	resp, err := HttpRequest{
		Url:    UrlFiles + "/" + fileId,
		ApiKey: apiKey,
	}.Delete()
	if err != nil {
		return fmt.Errorf("%w; error delete file api request", err)
	}
	var respObj struct {
		Object  string
		Deleted bool
	}
	if err := json.Unmarshal(resp, &respObj); err != nil {
		return fmt.Errorf("%w; error json unmarshalling delete file api response", err)
	}
	return nil
}
