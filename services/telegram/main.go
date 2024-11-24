package telegram

import (
	"fmt"
	"path/filepath"
	"time"

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

func (c *Telegram) DownloadFile(fileId int32) (string, error) {

	_, err := c.Client.DownloadFile(&client.DownloadFileRequest{
		FileId:   fileId,
		Priority: 32,
	})
	if err != nil {
		return "", fmt.Errorf("failed to download file with id - %v: %v", fileId, err)
	}

	for {
		file, err := c.Client.GetFile(&client.GetFileRequest{
			FileId: fileId,
		})
		if err != nil {
			return "", fmt.Errorf("failed to get file with id - %v on local machine while check file download status: %v", fileId, err)
		}

		if file.Local.IsDownloadingCompleted {
			return file.Local.Path, nil
		}

		time.Sleep(1 * time.Second)
	}
}
