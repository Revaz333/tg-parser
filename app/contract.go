package app

import "tg-lib/services/telegram"

type (
	Job struct {
		Content telegram.TGMessage
	}

	CarResponse struct {
		Brand        string  `json:"brand"`
		City         string  `json:"city"`
		Color        string  `json:"color"`
		DriveType    string  `json:"driveType"`
		EngineVolume float64 `json:"engineVolume"`
		FuelType     string  `json:"fuelType"`
		Mileage      int     `json:"mileage"`
		Model        string  `json:"model"`
		Price        int     `json:"price"`
		ReleaseYear  int     `json:"releaseYear"`
		Transmission string  `json:"transmission"`
		Phone        string  `json:"phone"`
		Description  string  `json:"description"`
	}

	FinalAdStruct struct {
		Info        CarResponse `json:"info"`
		Pictures    []Picture   `json:"pictures"`
		TGChannelID int64       `json:"tg_channel_id"`
	}

	Picture struct {
		Date  int64 `json:"date"`
		Paths Paths `json:"paths"`
	}

	Paths struct {
		UpName string `json:"upName"`
		Type   string `json:"type"`
		Path   string `json:"path"`
		Sizes  Sizes  `json:"sizes"`
	}

	Sizes struct {
		Small string `json:"small"`
	}

	NewMessageData struct {
		Messages    []telegram.TGMessage
		MessagesIDS []int64
	}
)

const (
	TypeMessage     = "updateNewMessage"
	TypeMessageText = "messageText"
)

var (
	channels = map[string]bool{
		"-1002392190527": true,
	}
)
