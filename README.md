# ChatGPT Golang Library

This is a Golang library for interacting with the OpenAI API.
It is meant to be imported into other projects.
`example.go` shows how to use the library.

## Example Usage

```sh
export OPENAI_API_KEY=ABC123
go run example.go models
go run example.go completion "who are you?"
go run example.go files
go run example.go upload test.jsonl
go run example.go delete file-xxx
go run example.go fine-tunes
go run example.go fine-tune-create file-xxx
go run example.go fine-tune-cancel ft-xxx
go run example.go fine-tune-events ft-xxx
go run example.go image "office with sunset view"
```

## Links

- Pricing: https://openai.com/api/pricing/
- API Keys: https://beta.openai.com/account/api-keys
- API Docs: https://beta.openai.com/docs/api-reference/introduction
- Models: https://beta.openai.com/docs/models/overview
