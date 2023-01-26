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
		prompt := os.Args[2]
		completion, err := api.GetCompletion(apiKey, prompt)
		if err != nil {
			exit1(fmt.Errorf("%s; error getting completion", err).Error())
		}
		for _, choice := range completion.Choices {
			fmt.Printf("Completion choice %d: %s\n", choice.Index, strings.TrimSpace(choice.Text))
		}
	default:
		exit1(fmt.Sprintf("Unknown command: %s", os.Args[1]))
	}
}

func exit1(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
