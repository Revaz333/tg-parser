package db

type (
	NewAdParams struct {
		MarkID         int64  `json:"mark_id"`
		ModelID        int64  `json:"model_id"`
		CityID         int64  `json:"city_id"`
		Images         []byte `json:"images"`
		Price          int    `json:"price"`
		Mileage        int    `json:"mileage"`
		DriveTypeID    int64  `json:"drive_type_id"`
		ReleaseYear    int    `json:"release_year"`
		TransmissionID int64  `json:"transmission_id"`
		ColorID        int64  `json:"color_id"`
		SourceType     string `json:"source_type"` // "tg_group" or "original"
		TGChannelID    int64  `json:"tg_channel_id,omitempty"`
		FuelTypeID     int64  `json:"fuel_type_id,omitempty"`
		EngineVolumeID int64  `json:"engine_volume_id,omitempty"`
	}
)
