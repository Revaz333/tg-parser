package main

import (
	"fmt"
	"tg-lib/app"
	"tg-lib/db"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	db, err := db.NewClient()
	if err != nil {
		log.Errorf("failed to init db connection: %v", err)
	}

	config, err := loadConfig()
	if err != nil {
		log.Errorf("failed to load config data fromt file: %v", err)
	}

	app := app.Init(config, db)
	app.StartAndLoop()
}

func loadConfig() (app.Config, error) {

	// set our config path
	viper.SetConfigFile("config/config.yml")

	// set config default values
	setDefaults()

	// read config to buffer
	if err := viper.ReadInConfig(); err != nil {
		return app.Config{}, fmt.Errorf("failed to read config: %v", err)
	}

	// bing config data to our struct
	var config app.Config
	if err := viper.Unmarshal(&config); err != nil {
		return app.Config{}, fmt.Errorf("failed to unpack config: %v", err)
	}

	return config, nil
}

func setDefaults() {
	viper.SetDefault("App.ID", "")
	viper.SetDefault("App.Hash", "")
}
