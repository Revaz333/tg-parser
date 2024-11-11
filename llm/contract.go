package llm

type (
	ChatMessage struct {
		ModelURI          string            `json:"modelUri"`
		CompletionOptions CompletionOptions `json:"completionOptions"`
		Messages          []Messages        `json:"messages"`
	}

	CompletionOptions struct {
		Stream      bool    `json:"stream"`
		Temperature float64 `json:"temperature"`
		MaxTokens   string  `json:"maxTokens"`
	}

	Messages struct {
		Role string `json:"role"`
		Text string `json:"text"`
	}

	ChatMessageResponse struct {
		Result Result `json:"result"`
	}
	Message struct {
		Role string `json:"role"`
		Text string `json:"text"`
	}
	Alternatives struct {
		Message Message `json:"message"`
		Status  string  `json:"status"`
	}
	Usage struct {
		InputTextTokens  string `json:"inputTextTokens"`
		CompletionTokens string `json:"completionTokens"`
		TotalTokens      string `json:"totalTokens"`
	}
	Result struct {
		Alternatives []Alternatives `json:"alternatives"`
		Usage        Usage          `json:"usage"`
		ModelVersion string         `json:"modelVersion"`
	}
)
