package telegram

import (
	"fmt"
	"path/filepath"

	"github.com/zelenin/go-tdlib/client"
)

type Telegram struct {
	Client *client.Client
}

func NewClient(appId int32, appToken string) (*Telegram, error) {
	authorizer := client.ClientAuthorizer()
	go client.CliInteractor(authorizer)

	authorizer.TdlibParameters <- &client.SetTdlibParametersRequest{
		UseTestDc:           false,
		DatabaseDirectory:   filepath.Join(".tdlib", "database"),
		FilesDirectory:      filepath.Join(".tdlib", "files"),
		UseFileDatabase:     true,
		UseChatInfoDatabase: true,
		UseMessageDatabase:  true,
		UseSecretChats:      false,
		ApiId:               appId,
		ApiHash:             appToken,
		SystemLanguageCode:  "en",
		DeviceModel:         "Server",
		SystemVersion:       "1.0.0",
		ApplicationVersion:  "1.0.0",
	}

	_, err := client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 1,
	})
	if err != nil {
		return &Telegram{}, fmt.Errorf("failed to set log level: %v", err)
	}

	tdlibClient, err := client.NewClient(authorizer)
	if err != nil {
		return &Telegram{}, fmt.Errorf("failed to create new client: %v", err)
	}

	return &Telegram{Client: tdlibClient}, nil
}
