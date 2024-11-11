package llm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

type LLM struct {
	client *resty.Client
}

const (
	APIhost   = "https://llm.api.cloud.yandex.net/foundationModels/v1"
	BasePoint = "/completion"
)

func NewClient(apiKey string) (*LLM, error) {

	client := resty.New()
	client.SetBaseURL(APIhost)
	client.SetAuthScheme("Api-key")
	client.SetAuthToken(apiKey)

	return &LLM{client}, nil
}

func (llm LLM) Send(messages Messages) (ChatMessageResponse, error) {

	promptBytes, err := os.ReadFile("config/prompt.json")
	if err != nil {
		return ChatMessageResponse{}, fmt.Errorf("failed to read prompt: %v", err)
	}

	var chatCompilation ChatMessage

	err = json.Unmarshal(promptBytes, &chatCompilation)
	if err != nil {
		return ChatMessageResponse{}, fmt.Errorf("failed to decode prompt bytes: %v", err)
	}

	chatCompilation.Messages = append(chatCompilation.Messages, messages)

	resp, err := llm.client.R().
		SetBody(chatCompilation).
		Post(BasePoint)

	if err != nil {
		return ChatMessageResponse{}, fmt.Errorf("an error occured while send request to LLM: %v", resp.Error())
	}

	if resp.StatusCode() != http.StatusOK {
		return ChatMessageResponse{}, fmt.Errorf("LLM resturns not %v code: %v", http.StatusOK, resp.Error())
	}

	var response ChatMessageResponse

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return ChatMessageResponse{}, fmt.Errorf("failed to bind response from LLM: %v", err)
	}

	return response, nil
}
