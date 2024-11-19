package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"tg-lib/db"
	"tg-lib/services/llm"
	"tg-lib/services/telegram"
	"time"

	log "github.com/sirupsen/logrus"
)

func (a App) ProcessMessage(messages []telegram.TGMessage) {

	var (
		finalResult FinalAdStruct
		err         error
	)

	for _, msg := range messages {

		finalResult.TGChannelID = msg.Message.ChatID

		if msg.Message.Content.Caption.Text != "" {
			log.Info("send request to llm")

			finalResult.Info, err = a.getAdInfo(msg.Message.Content.Caption.Text)
			if err != nil {
				log.Errorf("failed to get ad info from llm: %v", err)
				return
			}

			log.Info("request to llm done")
		}

		if len(msg.Message.Content.Photo.Sizes) != 0 {
			log.Info("get pictures")

			pictureId := msg.Message.Content.Photo.Sizes[len(msg.Message.Content.Photo.Sizes)-1].Photo.ID
			picture, err := a.getPicture(pictureId)
			if err != nil {
				log.Errorf("failed to get message picture with id - %v: %v", pictureId, err)
				continue
			}

			finalResult.Pictures = append(finalResult.Pictures, picture)

			log.Info("get pictures action done")
		}
	}

	ad, err := a.collectAd(finalResult)
	if err != nil {
		log.Errorf("failed to collect new ad params: %v", err)
		return
	}

	err = a.DB.CreateAd(ad)
	if err != nil {
		log.Errorf("failed to create new ad: %v", err)
		return
	}
}

func (a App) getAdInfo(text string) (CarResponse, error) {
	response, err := a.LLM.Send(
		llm.Messages{
			Role: "user",
			Text: text,
		},
	)
	if err != nil {
		return CarResponse{}, fmt.Errorf("failed to get response from llm: %v", err)
	}

	var result CarResponse
	resultString := strings.ReplaceAll(response.Result.Alternatives[0].Message.Text, "`", "")

	err = json.Unmarshal([]byte(resultString), &result)
	if err != nil {
		return CarResponse{}, fmt.Errorf("failed to decode llm response body: %v", err)
	}

	return result, nil
}

func (a App) getPicture(pictureId int) (Picture, error) {

	filePath, err := a.Tdlib.DownloadFile(int32(pictureId))
	if err != nil {
		return Picture{}, fmt.Errorf("failed to get file from telegram: %v", err)
	}

	fmt.Println("dddd pict path", filePath)
	// open downloaded file
	file, err := os.Open(filePath)
	if err != nil {
		return Picture{}, fmt.Errorf("failed to open file by path - %s: %v", filePath, err)
	}
	defer file.Close()

	// create dest file
	destFileName := fmt.Sprintf("telegram_image_%v%s", time.Now().UnixMicro(), filepath.Ext(filePath))
	destFilePath := fmt.Sprintf("%s/%s", a.Config.App.StoragePath, destFileName)

	destFile, err := os.Create(destFilePath)
	if err != nil {
		return Picture{}, fmt.Errorf("failed to create file by path - %s: %v", destFilePath, err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return Picture{}, fmt.Errorf("failed to copy file data from - %s to %s: %v", filePath, destFilePath, err)
	}

	// remove old tile from telegram folder
	err = os.Remove(filePath)
	if err != nil {
		return Picture{}, fmt.Errorf("failed to remote file with path - %s: %v", filePath, err)
	}

	return Picture{
		UpName: destFileName,
		Type:   "image",
		Path:   destFilePath,
		Sizes: Sizes{
			Small: destFileName,
		},
	}, nil
}

func (a App) collectAd(args FinalAdStruct) (db.NewAdParams, error) {

	var (
		err error
	)

	imageBytes, err := json.Marshal(args.Pictures)
	if err != nil {
		return db.NewAdParams{}, fmt.Errorf("failed to marshal pictures data: %v", err)
	}

	mark, err := a.DB.FindORCreateMark(args.Info.Brand)
	if err != nil {
		return db.NewAdParams{}, fmt.Errorf("failed to get mark: %v", err)
	}

	model, err := a.DB.FindORCreateModel(args.Info.Model, mark.ID)
	if err != nil {
		return db.NewAdParams{}, fmt.Errorf("failed to get model: %v", err)
	}

	city, err := a.DB.FindORCreateCity(args.Info.City)
	if err != nil {
		return db.NewAdParams{}, fmt.Errorf("failed to get city: %v", err)
	}

	driveType, err := a.DB.FindORCreateCity(args.Info.DriveType)
	if err != nil {
		return db.NewAdParams{}, fmt.Errorf("failed to get driveType: %v", err)
	}

	transmission, err := a.DB.FindORCreateTransmission(args.Info.Transmission)
	if err != nil {
		return db.NewAdParams{}, fmt.Errorf("failed to get transmission: %v", err)
	}

	fuelType, err := a.DB.FindORCreateFuelType(args.Info.FuelType)
	if err != nil {
		return db.NewAdParams{}, fmt.Errorf("failed to get fuelType: %v", err)
	}

	engineVolume, err := a.DB.FindORCreateEngineVolume(args.Info.EngineVolume)
	if err != nil {
		return db.NewAdParams{}, fmt.Errorf("failed to get engineVolume: %v", err)
	}

	tgChannel, err := a.DB.FindTgChannel(args.TGChannelID)
	if err != nil {
		return db.NewAdParams{}, fmt.Errorf("failed to get engineVolume: %v", err)
	}

	return db.NewAdParams{
		MarkID:         mark.ID,
		ModelID:        model.ID,
		CityID:         city.ID,
		DriveTypeID:    driveType.ID,
		TransmissionID: transmission.ID,
		FuelTypeID:     fuelType.ID,
		EngineVolumeID: engineVolume.ID,
		Images:         imageBytes,
		SourceType:     "tg_group",
		ColorID:        1,
		TGChannelID:    tgChannel.ID,
	}, nil
}
