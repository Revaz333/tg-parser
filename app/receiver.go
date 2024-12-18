package app

import (
	"encoding/json"
	"tg-lib/services/telegram"
	"time"

	log "github.com/sirupsen/logrus"
)

func (a *App) NewReceiver() error {

	// a.Tdlib
	listener := a.Tdlib.Client.GetListener()
	defer listener.Close()

	fullmessageMap := make(map[string][]telegram.TGMessage)
	messageTimer := make(map[string]*time.Timer)
	// listener.

	for upd := range listener.Updates {

		if upd.GetType() != TypeMessage {
			continue
		}

		var message telegram.TGMessage

		updBytes, err := json.Marshal(upd)
		if err != nil {
			log.Errorf("failed to convert update data to bytes: %v", err)
			continue
		}

		err = json.Unmarshal(updBytes, &message)
		if err != nil {
			log.Errorf("failed to decode update data bytes: %v", err)
			continue
		}

		// if _, ok := channels[fmt.Sprintf("%v", message.Message.ChatID)]; !ok {
		// 	continue
		// }

		fullmessageMap[message.Message.MediaAlbumID] = append(fullmessageMap[message.Message.MediaAlbumID], message)
		if timer, exists := messageTimer[message.Message.MediaAlbumID]; exists {
			timer.Stop()
		}

		messageTimer[message.Message.MediaAlbumID] = time.AfterFunc(1*time.Second, func() {
			go a.ProcessMessage(fullmessageMap[message.Message.MediaAlbumID])
		})
	}

	return nil
}
