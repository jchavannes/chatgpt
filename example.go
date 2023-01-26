package main

import (
	"fmt"
	"github.com/jchavannes/chatgpt/api"
	"os"
	"time"
)

const EnvOpenAiKey = "OPENAI_API_KEY"

func main() {
	apiKey := os.Getenv(EnvOpenAiKey)
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run example.go <command>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "models":
		models, err := api.GetModelList(apiKey)
		if err != nil {
			panic(fmt.Errorf("%w; error getting model list", err))
		}
		fmt.Printf("Models: %d\n", len(models))
		for _, model := range models {
			fmt.Printf("Model: %s - %s\n", model.Id, time.Unix(model.Created, 0).Format(time.RFC3339))
		}
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}
