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
			exit1(fmt.Errorf("%s; error getting model list", err).Error())
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
			exit1(fmt.Errorf("%s; error getting completion", err).Error())
		}
		for _, choice := range completion.Choices {
			fmt.Printf("Completion choice %d: %s\n", choice.Index, strings.TrimSpace(choice.Text))
		}
	case "files":
		fileList, err := api.FileList(apiKey)
		if err != nil {
			exit1(fmt.Errorf("%s; error getting file list", err).Error())
		}
		fmt.Printf("Files: %d\n", len(fileList))
		for _, file := range fileList {
			fmt.Printf("File: %s - %s (%d)\n", file.Id, file.Filename, file.Bytes)
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
			exit1(fmt.Errorf("%s; error api file upload", err).Error())
		}
		fmt.Printf("File: %s - %s\n", file.Id, file.Filename)
	case "delete":
		if len(os.Args) < 3 || os.Args[2] == "" {
			exit1("Usage: go run example.go delete <filename>")
		}
		filename := os.Args[2]
		if err := api.FileDelete(apiKey, filename); err != nil {
			exit1(fmt.Errorf("%s; error api file delete", err).Error())
		}
		fmt.Printf("File deleted: %s\n", filename)
	default:
		exit1(fmt.Sprintf("Unknown command: %s", os.Args[1]))
	}
}

func exit1(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
