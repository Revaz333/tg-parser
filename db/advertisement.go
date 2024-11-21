package db

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func (db DB) CreateAd(args NewAdParams) error {
	result := db.Client.Table("advertisements").Create(&args)
	if result.Error != nil {
		return fmt.Errorf("failed to create new ad: %v", result.Error)
	}

	return nil
}

func (db DB) FindORCreateMark(title string) (Mark, error) {

	var mark Mark

	result := db.Client.Table("marks").Where("title like ?", "%"+title+"%").First(&mark)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return Mark{}, fmt.Errorf("failed to get mark with title - %s: %v", title, result.Error)
	} else if result.Error == gorm.ErrRecordNotFound {
		mark.Title = title
		result := db.Client.Table("marks").Create(&mark)
		if result.Error != nil {
			return Mark{}, fmt.Errorf("failed to create new mark with title - %s: %v", title, result.Error)
		}
	}

	return mark, nil
}

func (db DB) FindORCreateModel(title string, markId uint) (Model, error) {

	var model Model

	result := db.Client.Table("models").Where("title like ? AND mark_id = ?", "%"+title+"%", markId).First(&model)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return Model{}, fmt.Errorf("failed to get model with title - %s: %v", title, result.Error)
	} else if result.Error == gorm.ErrRecordNotFound {
		model.Title = title
		model.MarkID = uint(markId)
		result := db.Client.Table("models").Create(&model)
		if result.Error != nil {
			return Model{}, fmt.Errorf("failed to create new model with title - %s: %v", title, result.Error)
		}
	}

	return model, nil
}

func (db DB) FindORCreateCity(title string) (City, error) {

	var city City

	result := db.Client.Table("cities").Where("title like ?", "%"+title+"%").First(&city)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return City{}, fmt.Errorf("failed to get city with title - %s: %v", title, result.Error)
	} else if result.Error == gorm.ErrRecordNotFound {
		city.Title = title
		result := db.Client.Table("cities").Create(&city)
		if result.Error != nil {
			return City{}, fmt.Errorf("failed to create new city with title - %s: %v", title, result.Error)
		}
	}

	return city, nil
}

func (db DB) FindORCreateDriveType(title string) (DriveType, error) {

	var driveType DriveType

	result := db.Client.Table("drive_types").Where("title like ?", "%"+title+"%").First(&driveType)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return DriveType{}, fmt.Errorf("failed to get driveType with title - %s: %v", title, result.Error)
	} else if result.Error == gorm.ErrRecordNotFound {
		driveType.Title = title
		result := db.Client.Table("drive_types").Create(&driveType)
		if result.Error != nil {
			return DriveType{}, fmt.Errorf("failed to create new driveType with title - %s: %v", title, result.Error)
		}
	}

	return driveType, nil
}

func (db DB) FindORCreateTransmission(title string) (Transmission, error) {

	var transmission Transmission

	result := db.Client.Table("transmissions").Where("title like ?", "%"+title+"%").First(&transmission)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return Transmission{}, fmt.Errorf("failed to get transmission with title - %s: %v", title, result.Error)
	} else if result.Error == gorm.ErrRecordNotFound {
		transmission.Title = title
		result := db.Client.Table("transmissions").Create(&transmission)
		if result.Error != nil {
			return Transmission{}, fmt.Errorf("failed to create new transmission with title - %s: %v", title, result.Error)
		}
	}

	return transmission, nil
}

func (db DB) FindORCreateFuelType(title string) (FuelType, error) {

	var fuelType FuelType

	result := db.Client.Table("fuel_types").Where("title like ?", "%"+title+"%").First(&fuelType)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return FuelType{}, fmt.Errorf("failed to get fuelType with title - %s: %v", title, result.Error)
	} else if result.Error == gorm.ErrRecordNotFound {
		fuelType.Title = title
		result := db.Client.Table("fuel_types").Create(&fuelType)
		if result.Error != nil {
			return FuelType{}, fmt.Errorf("failed to create new fuelType with title - %s: %v", title, result.Error)
		}
	}

	return fuelType, nil
}

func (db DB) FindORCreateEngineVolume(volume float64) (EngineVolume, error) {

	var engineVolume EngineVolume

	result := db.Client.Table("engine_volumes").Where("volume = ?", volume).First(&engineVolume)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return EngineVolume{}, fmt.Errorf("failed to get engineVolume with volume - %v: %v", volume, result.Error)
	} else if result.Error == gorm.ErrRecordNotFound {
		engineVolume.Volume = volume
		result := db.Client.Table("engine_volumes").Create(&engineVolume)
		if result.Error != nil {
			return EngineVolume{}, fmt.Errorf("failed to create new engineVolume with volume - %v: %v", volume, result.Error)
		}
	}

	return engineVolume, nil
}

func (db DB) FindTgChannel(chatId int64) (TGChannel, error) {

	var (
		channel TGChannel
		chat    = strconv.Itoa(int(chatId))
	)

	result := db.Client.Table("tg_channels").Where("chat_id = ?", chat).First(&channel)
	if result.Error != nil {
		return TGChannel{}, fmt.Errorf("failed to get channel with chatId - %v: %v", chatId, result.Error)
	}

	return channel, nil
}
