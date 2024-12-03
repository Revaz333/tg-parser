package llm

type (
	ChatMessage struct {
		Model       string     `json:"model"`
		Messages    []Messages `json:"messages"`
		Temperature float64    `json:"temperature"`
		Stream      bool       `json:"stream"`
	}
	Messages struct {
		Role    string `json:"role"`
		Сontent string `json:"content"`
	}

	ChatMessageResponse struct {
		ID                string    `json:"id"`
		Object            string    `json:"object"`
		Created           int       `json:"created"`
		Model             string    `json:"model"`
		Choices           []Choices `json:"choices"`
		Usage             Usage     `json:"usage"`
		SystemFingerprint string    `json:"system_fingerprint"`
	}
	RespMessage struct {
		Role    string `json:"role"`
		Content string `json:"content"`
		Refusal any    `json:"refusal"`
	}
	Choices struct {
		Index        int         `json:"index"`
		Message      RespMessage `json:"message"`
		Logprobs     any         `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	}
	PromptTokensDetails struct {
		CachedTokens int `json:"cached_tokens"`
		AudioTokens  int `json:"audio_tokens"`
	}
	CompletionTokensDetails struct {
		ReasoningTokens          int `json:"reasoning_tokens"`
		AudioTokens              int `json:"audio_tokens"`
		AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
		RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
	}
	Usage struct {
		PromptTokens            int                     `json:"prompt_tokens"`
		CompletionTokens        int                     `json:"completion_tokens"`
		TotalTokens             int                     `json:"total_tokens"`
		PromptTokensDetails     PromptTokensDetails     `json:"prompt_tokens_details"`
		CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
	}
)
