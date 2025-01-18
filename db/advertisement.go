package db

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func (db DB) CreateAd(args NewAdParams) error {
	err := db.Client.Table("advertisements").Create(&args).Error
	if err != nil {
		return fmt.Errorf("failed to create new ad: %v", err)
	}

	err = db.EnableModeration(args.ID)
	if err != nil {
		return fmt.Errorf("moderation error: %v", err)
	}

	return nil
}

func (db DB) FindORCreateMark(title string) (Mark, error) {

	var mark Mark

	result := db.Client.Table("marks").Where("title like ?", "%"+title+"%").First(&mark)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return Mark{}, fmt.Errorf("failed to get mark with title - %s: %v", title, result.Error)
	} else if result.Error == gorm.ErrRecordNotFound {
		db.Cache.Delete("system_data")
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
		db.Cache.Delete("system_data")
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
		db.Cache.Delete("system_data")
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
		db.Cache.Delete("system_data")
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
		db.Cache.Delete("system_data")
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
		db.Cache.Delete("system_data")
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
		db.Cache.Delete("system_data")
		engineVolume.Volume = volume
		result := db.Client.Table("engine_volumes").Create(&engineVolume)
		if result.Error != nil {
			return EngineVolume{}, fmt.Errorf("failed to create new engineVolume with volume - %v: %v", volume, result.Error)
		}
	}

	return engineVolume, nil
}

func (db DB) FindOrCreateTgChannel(chatId int64) (TGChannel, error) {

	var (
		channel TGChannel
		chat    = strconv.Itoa(int(chatId))
	)

	result := db.Client.Table("tg_channels").Where("chat_id = ?", chat).First(&channel)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return TGChannel{}, fmt.Errorf("failed to get channel with chatId - %v: %v", chatId, result.Error)
	} else if result.Error == gorm.ErrRecordNotFound {
		db.Cache.Delete("system_data")
		channel.ChatID = chatId
		result := db.Client.Table("tg_channels").Create(&channel)
		if result.Error != nil {
			return TGChannel{}, fmt.Errorf("failed to create new teleram channel with tg_id - %v: %v", chatId, result.Error)
		}
	}

	return channel, nil
}

func (db DB) FindOrCreateColor(title string) (Color, error) {

	var color Color

	result := db.Client.Table("colors").Where("title = ?", title).First(&color)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return Color{}, fmt.Errorf("failed to get color with title - %s: %v", title, result.Error)
	} else if result.Error == gorm.ErrRecordNotFound {
		db.Cache.Delete("system_data")
		color.Title = title
		result := db.Client.Table("colors").Create(&color)
		if result.Error != nil {
			return Color{}, fmt.Errorf("failed to create new color with title - %s: %v", title, result.Error)
		}
	}

	return color, nil
}

func (db DB) GetMarksList() ([]string, error) {

	var marks []Mark
	result := db.Client.Table("marks").Find(&marks)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return []string{}, fmt.Errorf("failed to get marks list: %v", result.Error)
	}

	var resp []string
	for _, mark := range marks {
		resp = append(resp, mark.Title)
	}

	return resp, nil
}

func (db DB) GetModelsList() ([]string, error) {

	var models []Model
	result := db.Client.Table("models").Find(&models)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return []string{}, fmt.Errorf("failed to get models list: %v", result.Error)
	}

	var resp []string
	for _, model := range models {
		resp = append(resp, model.Title)
	}

	return resp, nil
}

func (db DB) GetColorsList() ([]string, error) {

	var colors []Color
	result := db.Client.Table("colors").Find(&colors)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return []string{}, fmt.Errorf("failed to get colors list: %v", result.Error)
	}

	var resp []string
	for _, color := range colors {
		resp = append(resp, color.Title)
	}

	return resp, nil
}

func (db DB) GetTransmissionsList() ([]string, error) {

	var transmissions []Transmission
	result := db.Client.Table("transmissions").Find(&transmissions)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return []string{}, fmt.Errorf("failed to get transmissions list: %v", result.Error)
	}

	var resp []string
	for _, transmission := range transmissions {
		resp = append(resp, transmission.Title)
	}

	return resp, nil
}

func (db DB) GetCitiesList() ([]string, error) {

	var cities []City
	result := db.Client.Table("cities").Find(&cities)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return []string{}, fmt.Errorf("failed to get cities list: %v", result.Error)
	}

	var resp []string
	for _, city := range cities {
		resp = append(resp, city.Title)
	}

	return resp, nil
}

func (db DB) GetEngineVolumesList() ([]float64, error) {

	var engineVolumes []EngineVolume
	result := db.Client.Table("engine_volumes").Find(&engineVolumes)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return []float64{}, fmt.Errorf("failed to get engine_volumes list: %v", result.Error)
	}

	var resp []float64
	for _, engineVolume := range engineVolumes {
		resp = append(resp, engineVolume.Volume)
	}

	return resp, nil
}

func (db DB) GetFuelTypesList() ([]string, error) {

	var fuelTypes []FuelType
	result := db.Client.Table("fuel_types").Find(&fuelTypes)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return []string{}, fmt.Errorf("failed to get fuelTypes list: %v", result.Error)
	}

	var resp []string
	for _, fuelType := range fuelTypes {
		resp = append(resp, fuelType.Title)
	}

	return resp, nil
}

func (db DB) GetDriveTypesList() ([]string, error) {

	var driveTypes []FuelType
	result := db.Client.Table("drive_types").Find(&driveTypes)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return []string{}, fmt.Errorf("failed to get drive_types list: %v", result.Error)
	}

	var resp []string
	for _, driveType := range driveTypes {
		resp = append(resp, driveType.Title)
	}

	return resp, nil
}

func (db DB) GetData() (map[string]interface{}, error) {

	// data, found := db.Cache.Get("system_data")
	// if found {
	// 	return data.(map[string]interface{}), nil
	// }

	data, err := db.CollectData()
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("collext data error: %v", err)
	}

	// db.Cache.Set("system_data", data, 24*time.Hour)

	return data, nil
}

func (db DB) EnableModeration(adId uint) error {

	m := Moderation{
		AdID: adId,
	}

	err := db.Client.Create(&m).Error
	if err != nil {
		return fmt.Errorf("failed to enable moderation on ad with id - %v: %v", adId, err)
	}

	return nil
}

func (db DB) CollectData() (map[string]interface{}, error) {

	var (
		err  error
		resp = map[string]interface{}{}
	)

	resp["colors"], err = db.GetColorsList()
	if err != nil {
		return map[string]interface{}{}, err
	}

	resp["marks"], err = db.GetMarksList()
	if err != nil {
		return map[string]interface{}{}, err
	}

	resp["models"], err = db.GetModelsList()
	if err != nil {
		return map[string]interface{}{}, err
	}

	resp["cities"], err = db.GetCitiesList()
	if err != nil {
		return map[string]interface{}{}, err
	}

	resp["engine_volumes"], err = db.GetEngineVolumesList()
	if err != nil {
		return map[string]interface{}{}, err
	}

	resp["drive_types"], err = db.GetDriveTypesList()
	if err != nil {
		return map[string]interface{}{}, err
	}

	resp["fuel_types"], err = db.GetFuelTypesList()
	if err != nil {
		return map[string]interface{}{}, err
	}

	resp["transmissions"], err = db.GetTransmissionsList()
	if err != nil {
		return map[string]interface{}{}, err
	}

	return resp, nil
}
