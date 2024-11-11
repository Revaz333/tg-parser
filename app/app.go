package app

import (
	"encoding/json"
	"fmt"
	"tg-lib/db"
	"tg-lib/llm"
	"tg-lib/telegram"

	log "github.com/sirupsen/logrus"
)

var (
	channels = map[string]bool{
		"-1002392190527": true,
	}

	TypeMessage = "updateNewMessage"
)

type (
	App struct {
		Config Config
		DB     *db.DB
	}

	Config struct {
		App struct {
			ID   int32
			Hash string
		}
		LLM struct {
			Key string
		}
	}

	CarResponse struct {
		Brand        string  `json:"brand"`
		City         string  `json:"city"`
		Color        any     `json:"color"`
		DriveType    any     `json:"driveType"`
		EngineVolume float64 `json:"engineVolume"`
		FuelType     any     `json:"fuelType"`
		Mileage      int     `json:"mileage"`
		Model        string  `json:"model"`
		Price        int     `json:"price"`
		ReleaseYear  int     `json:"releaseYear"`
		Transmission string  `json:"transmission"`
	}
)

func Init(config Config, db *db.DB) *App {
	return &App{config, db}
}

func (a App) StartAndLoop() error {

	client, err := telegram.NewClient(a.Config.App.ID, a.Config.App.Hash)
	if err != nil {
		log.Errorf("failed to init new telegram client: %v", err)
		return fmt.Errorf("failed to init new telegram client: %v", err)
	}

	llmClient, err := llm.NewClient(a.Config.LLM.Key)
	if err != nil {
		log.Errorf("failed to init new LLM client: %v", err)
		return fmt.Errorf("failed to init new LLM client: %v", err)
	}

	updates := client.Client.GetListener().Updates
	for upd := range updates {
		if upd.GetType() != TypeMessage {
			continue
		}

		var message telegram.TGMessage
		b, _ := json.Marshal(upd)
		json.Unmarshal(b, &message)

		if _, ok := channels[fmt.Sprintf("%v", message.Message.SenderID.ChatID)]; !ok {
			continue
		}

		log.Info("New request to LLM")

		reqData := llm.Messages{
			Role: "user",
			Text: message.Message.Content.Text.Text,
		}

		resp, err := llmClient.Send(reqData)
		if err != nil {
			log.Errorf("failed to get response from LLM: %v", err)
			continue
		}

		log.Info("end request to LLM")

		var car CarResponse

		err = json.Unmarshal([]byte(resp.Result.Alternatives[0].Message.Text), &car)
		if err != nil {
			log.Errorf("failed to get car data from LLM json response: %v", err)
			continue
		}

		fmt.Printf("car models is: %s", car.Model)
	}

	return nil
}
