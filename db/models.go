package db

type (
	NewAdParams struct {
		ID             uint   `gorm:"primaryKey"`
		MarkID         uint   `json:"mark_id"`
		ModelID        uint   `json:"model_id"`
		CityID         uint   `json:"city_id"`
		Images         []byte `json:"images"`
		Price          int    `json:"price"`
		Mileage        int    `json:"mileage"`
		DriveTypeID    uint   `json:"drive_type_id"`
		ReleaseYear    int    `json:"release_year"`
		TransmissionID uint   `json:"transmission_id"`
		ColorID        uint   `json:"color_id"`
		SourceType     string `json:"source_type"` // "tg_group" or "original"
		TGChannelID    uint   `json:"tg_channel_id,omitempty"`
		FuelTypeID     uint   `json:"fuel_type_id,omitempty"`
		EngineVolumeID uint   `json:"engine_volume_id,omitempty"`
	}

	Mark struct {
		ID    uint   `gorm:"primaryKey"`
		Title string `json:"title"`
	}

	Model struct {
		ID     uint   `gorm:"primaryKey"`
		Title  string `json:"title"`
		MarkID uint   `json:"mark_id"`
		Mark   Mark   `gorm:"foreignKey:MarkID"`
	}

	City struct {
		ID    uint   `gorm:"primaryKey"`
		Title string `json:"title"`
	}

	DriveType struct {
		ID    uint   `gorm:"primaryKey"`
		Title string `json:"title"`
	}

	Transmission struct {
		ID    uint   `gorm:"primaryKey"`
		Title string `json:"title"`
	}

	FuelType struct {
		ID    uint   `gorm:"primaryKey"`
		Title string `json:"title"`
	}

	EngineVolume struct {
		ID     uint    `gorm:"primaryKey"`
		Volume float64 `json:"volume"`
	}

	TGChannel struct {
		ID     uint   `gorm:"primaryKey"`
		ChatID string `json:"chat_id"`
	}
)
