package app

import "tg-lib/services/telegram"

type (
	Job struct {
		Content telegram.TGMessage
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
		Phone        string  `json:"phone"`
	}

	FinalAdStruct struct {
		Info     CarResponse `json:"info"`
		Pictures []Picture   `json:"pictures"`
	}

	Picture struct {
		UpName string `json:"upName"`
		Type   string `json:"type"`
		Path   string `json:"path"`
		Sizes  Sizes  `json:"sizes"`
	}
	Sizes struct {
		Small string `json:"small"`
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
