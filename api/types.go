package api

import (
	"fmt"
	"time"
)

type baseObject struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
}

type Completion struct {
	baseObject
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param"`
	Type    string `json:"type"`
}

type Event struct {
	baseObject
	Level   string `json:"level"`
	Message string `json:"message"`
}

func (e Event) Info() string {
	return fmt.Sprintf("Event: %s %s %s", e.Object, e.Level, e.Message)
}

type File struct {
	baseObject
	Bytes    int64  `json:"bytes"`
	Filename string `json:"filename"`
	Purpose  string `json:"purpose"`
}

func (f File) Info() string {
	return fmt.Sprintf("File: %s - %s (%d)", f.Id, f.Filename, f.Bytes)
}

type FineTune struct {
	baseObject
	Model           string      `json:"model"`
	FineTunedModel  string      `json:"fine_tuned_model"`
	HyperParams     interface{} `json:"hyperparams"`
	OrganizationId  string      `json:"organization_id"`
	ResultFiles     []File      `json:"result_files"`
	Status          string      `json:"status"`
	ValidationFiles []File      `json:"validation_files"`
	TrainingFiles   []File      `json:"training_files"`
	UpdatedAt       int64       `json:"updated_at"`
	Events          []Event     `json:"events"`
}

func (f FineTune) Info() string {
	return fmt.Sprintf("Fine tune: %s - %s %s %s",
		f.Id, f.Model, f.Status, time.Unix(f.UpdatedAt, 0).Format(time.RFC3339))
}

type Image struct {
	Url string `json:"url"`
}

type Model struct {
	baseObject
	OwnedBy    string       `json:"owned_by"`
	Permission []Permission `json:"permission"`
	Root       string       `json:"root"`
	Parent     string       `json:"parent"`
}

func (m Model) Info() string {
	var parent string
	if m.Parent != "" {
		parent = " (parent: " + m.Parent + ")"
	}
	return fmt.Sprintf("Model: %s - %s%s", m.Id, time.Unix(m.Created, 0).Format(time.RFC3339), parent)
}

type Permission struct {
	baseObject
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	Organization       string      `json:"organization"`
	Group              interface{} `json:"group"`
	IsBlocking         bool        `json:"is_blocking"`
}

type CompletionRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Temperature float64 `json:"temperature"`
	MaxTokens   int     `json:"max_tokens"`
}

type FineTuneCreateRequest struct {
	TrainingFile string `json:"training_file"`
}

type ImageRequest struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}
