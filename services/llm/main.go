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
	apiKey string
}

const (
	BasePoint = "/send-message"
)

func NewClient(apiHost, apiKey string) (*LLM, error) {

	client := resty.New()
	client.SetBaseURL(apiHost)
	// client.SetAuthScheme("Bearer")
	// client.SetAuthToken(apiKey)
	// client.SetHeader("Autorization", "Bearer "+apiKey)

	return &LLM{client, apiKey}, nil
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
		SetHeader("Authorization", llm.apiKey).
		SetBody(chatCompilation).
		Post(BasePoint)

	if err != nil {
		return ChatMessageResponse{}, fmt.Errorf("an error occured while send request to LLM: %v", resp.Error())
	}

	if resp.StatusCode() != http.StatusOK {
		return ChatMessageResponse{}, fmt.Errorf("LLM resturns not %v code: %v", http.StatusOK, resp)
	}

	var response ChatMessageResponse

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return ChatMessageResponse{}, fmt.Errorf("failed to bind response from LLM: %v", err)
	}

	return response, nil
}
