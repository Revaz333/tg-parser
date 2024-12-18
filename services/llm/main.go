package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
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

	return &LLM{client, apiKey}, nil
}

func (llm LLM) Send(messages Messages, vars map[string]interface{}) (ChatMessageResponse, error) {

	promptBytes, err := os.ReadFile("config/prompt.json")
	if err != nil {
		return ChatMessageResponse{}, fmt.Errorf("failed to read prompt: %v", err)
	}

	tpl, err := template.New("message_data").Parse(string(promptBytes))
	if err != nil {
		return ChatMessageResponse{}, fmt.Errorf("an error occure while parse prompt template: %v", err)
	}

	var tplBuff bytes.Buffer

	tpl.Execute(&tplBuff, vars)

	var chatCompilation ChatMessage

	err = json.Unmarshal(tplBuff.Bytes(), &chatCompilation)
	if err != nil {
		return ChatMessageResponse{}, fmt.Errorf("failed to decode prompt bytes: %v", err)
	}

	chatCompilation.Messages = append(chatCompilation.Messages, messages)
	resp, err := llm.client.R().
		SetHeader("Authorization", "Bearer "+llm.apiKey).
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
