package app

import (
	"fmt"
	"tg-lib/db"
	"tg-lib/services/llm"
	"tg-lib/services/telegram"

	log "github.com/sirupsen/logrus"
)

type (
	DBInterface interface {
		CreateAd(args db.NewAdParams) error
	}
	App struct {
		Config Config
		DB     DBInterface
		LLM    *llm.LLM
		Tdlib  *telegram.Telegram
	}

	Config struct {
		App struct {
			ID          int32
			Hash        string
			StoragePath string
		}
		LLM struct {
			Key string
		}
		DB struct {
			Dsn string
		}
	}
)

func Init(config Config) *App {

	db, err := db.NewClient(config.DB.Dsn)
	if err != nil {
		log.Errorf("failed to init db connection: %v", err)
	}

	llm, _ := llm.NewClient(config.LLM.Key)
	tglib, err := telegram.NewClient(config.App.ID, config.App.Hash)
	if err != nil {
		log.Errorf("failed to init new telegram client: %v", err)
	}

	return &App{Config: config, LLM: llm, Tdlib: tglib, DB: db}
}

func (a App) StartAndLoop() error {

	err := a.NewReceiver()
	if err != nil {
		return fmt.Errorf("failed to init jobs reciver: %v", err)
	}

	return nil
}
