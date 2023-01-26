package api

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

type File struct {
	baseObject
	Bytes    int64  `json:"bytes"`
	Filename string `json:"filename"`
	Purpose  string `json:"purpose"`
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
}

type Model struct {
	baseObject
	OwnedBy    string       `json:"owned_by"`
	Permission []Permission `json:"permission"`
	Root       string       `json:"root"`
	Parent     *Model       `json:"parent"`
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
