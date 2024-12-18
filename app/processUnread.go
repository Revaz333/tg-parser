package app

import (
	"encoding/json"
	"tg-lib/services/telegram"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/zelenin/go-tdlib/client"
)

func (a App) ProcessUnreadMessages() {

	chats, err := a.Tdlib.Client.GetChats(&client.GetChatsRequest{
		Limit: 1000,
	})
	if err != nil {
		log.Errorf("failed to get client chats: %v", err)
		return
	}

	for {
		for _, chatID := range chats.ChatIds {
			chat, err := a.Tdlib.Client.GetChat(&client.GetChatRequest{
				ChatId: chatID,
			})
			if err != nil {
				log.Errorf("failed to get chat by id - %v: %v", chatID, err)
				continue
			}

			if chat.UnreadCount == 0 {
				continue
			}

			chatHistory, err := a.Tdlib.Client.GetChatHistory(&client.GetChatHistoryRequest{
				ChatId:        chatID,
				Limit:         chat.UnreadCount,
				FromMessageId: 0,
			})
			if err != nil {
				log.Errorf("failed to get history from chat by id - %v: %v", chatID, err)
				continue
			}

			err = a.reciveUnreadMessages(chatID, chatHistory.Messages)
			if err != nil {
				log.Errorf("failed to recive history from chat by id - %v: %v", chatID, err)
				continue
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func (a App) reciveUnreadMessages(chatId int64, messages []*client.Message) error {

	fullmessageMap := make(map[string][]telegram.TGMessage)
	messageTimer := make(map[string]*time.Timer)

	for _, msg := range messages {
		var message telegram.Message

		updBytes, err := json.Marshal(msg)
		if err != nil {
			log.Errorf("failed to convert update data to bytes: %v", err)
			continue
		}

		err = json.Unmarshal(updBytes, &message)
		if err != nil {
			log.Errorf("failed to decode update data bytes: %v", err)
			continue
		}

		fullmessageMap[string(message.MediaAlbumID)] = append(fullmessageMap[string(message.MediaAlbumID)], telegram.TGMessage{Message: message})
		if timer, exists := messageTimer[string(message.MediaAlbumID)]; exists {
			timer.Stop()
		}

		messageTimer[string(message.MediaAlbumID)] = time.AfterFunc(1*time.Second, func() {

			var readedMsgs []int64

			for _, msg := range fullmessageMap[string(message.MediaAlbumID)] {
				readedMsgs = append(readedMsgs, msg.Message.ID)
			}

			_, err := a.Tdlib.Client.ViewMessages(&client.ViewMessagesRequest{
				ChatId:     chatId,
				MessageIds: readedMsgs,
				ForceRead:  true,
			})
			if err != nil {
				log.Errorf("failed to read message: %v", err)
				return
			}

			go a.ProcessMessage(fullmessageMap[string(message.MediaAlbumID)])
		})
	}

	return nil
}
