package main

import (
	"fmt"
	"github.com/jchavannes/chatgpt/api"
	"os"
	"strings"
	"time"
)

const EnvOpenAiKey = "OPENAI_API_KEY"

func main() {
	apiKey := os.Getenv(EnvOpenAiKey)
	if apiKey == "" {
		exit1(fmt.Sprintf("Missing %s environment variable", EnvOpenAiKey))
	}
	if len(os.Args) < 2 {
		exit1("Usage: go run example.go <command>")
	}
	switch os.Args[1] {
	case "models":
		models, err := api.GetModelList(apiKey)
		if err != nil {
			exit1(fmt.Errorf("%w; error getting model list", err).Error())
		}
		fmt.Printf("Models: %d\n", len(models))
		for _, model := range models {
			fmt.Printf("Model: %s - %s\n", model.Id, time.Unix(model.Created, 0).Format(time.RFC3339))
		}
	case "completion":
		if len(os.Args) < 3 {
			exit1("Usage: go run example.go completion <prompt>")
		}
		prompt := strings.Join(os.Args[2:], " ")
		completion, err := api.GetCompletion(apiKey, prompt)
		if err != nil {
			exit1(fmt.Errorf("%w; error getting completion", err).Error())
		}
		for _, choice := range completion.Choices {
			fmt.Printf("Completion choice %d: %s\n", choice.Index, strings.TrimSpace(choice.Text))
		}
	case "files":
		fileList, err := api.FileList(apiKey)
		if err != nil {
			exit1(fmt.Errorf("%w; error getting file list", err).Error())
		}
		fmt.Printf("Files: %d\n", len(fileList))
		for _, file := range fileList {
			fmt.Println(file.Info())
		}
	case "upload":
		if len(os.Args) < 3 || os.Args[2] == "" {
			exit1("Usage: go run example.go upload <filepath>")
		}
		filename := os.Args[2]
		if len(filename) < 7 || filename[len(filename)-6:] != ".jsonl" {
			exit1("error file name must end with .jsonl")
		}
		file, err := api.FileUpload(apiKey, filename)
		if err != nil {
			exit1(fmt.Errorf("%w; error api file upload", err).Error())
		}
		fmt.Println(file.Info())
	case "delete":
		if len(os.Args) < 3 || os.Args[2] == "" {
			exit1("Usage: go run example.go delete <filename>")
		}
		filename := os.Args[2]
		if err := api.FileDelete(apiKey, filename); err != nil {
			exit1(fmt.Errorf("%w; error api file delete", err).Error())
		}
		fmt.Printf("File deleted: %s\n", filename)
	case "fine-tunes":
		fineTunes, err := api.FineTuneList(apiKey)
		if err != nil {
			exit1(fmt.Errorf("%w; error getting fine tunes", err).Error())
		}
		fmt.Printf("Fine Tunes: %d\n", len(fineTunes))
		for _, fineTune := range fineTunes {
			fmt.Println(fineTune.Info())
		}
	case "fine-tune-create":
		if len(os.Args) < 3 || os.Args[2] == "" {
			exit1("Usage: go run example.go fine-tune-create <filename>")
		}
		filename := os.Args[2]
		fineTune, err := api.FineTuneCreate(apiKey, filename)
		if err != nil {
			exit1(fmt.Errorf("%w; error getting fine tunes", err).Error())
		}
		fmt.Println(fineTune.Info())
		for _, event := range fineTune.Events {
			fmt.Println(event.Info())
		}
	case "fine-tune-cancel":
		if len(os.Args) < 3 || os.Args[2] == "" {
			exit1("Usage: go run example.go fine-tune-cancel <fine_tune_id>")
		}
		fineTuneId := os.Args[2]
		fineTune, err := api.FineTuneCancel(apiKey, fineTuneId)
		if err != nil {
			exit1(fmt.Errorf("%w; error api file delete", err).Error())
		}
		fmt.Println(fineTune.Info())
	case "fine-tune-events":
		if len(os.Args) < 3 || os.Args[2] == "" {
			exit1("Usage: go run example.go fine-tune-events <fine_tune_id>")
		}
		fineTuneId := os.Args[2]
		events, err := api.FineTuneEvents(apiKey, fineTuneId)
		if err != nil {
			exit1(fmt.Errorf("%w; error getting fine tune events", err).Error())
		}
		fmt.Printf("Fine Tune Events: %d\n", len(events))
		for _, event := range events {
			fmt.Println(event.Info())
		}
	case "image":
		if len(os.Args) < 3 {
			exit1("Usage: go run example.go image <prompt>")
		}
		prompt := strings.Join(os.Args[2:], " ")
		image, err := api.ImageCreate(apiKey, prompt)
		if err != nil {
			exit1(fmt.Errorf("%w; error getting image create", err).Error())
		}
		fmt.Printf("Image: %s\n", image.Url)
	default:
		exit1(fmt.Sprintf("Unknown command: %s", os.Args[1]))
	}
}

func exit1(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
