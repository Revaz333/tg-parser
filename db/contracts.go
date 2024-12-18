package db

type (
	GetDataResponse struct {
		Cities        []string  `json:"cities"`
		Colors        []string  `json:"colors"`
		Marks         []string  `json:"marks"`
		Models        []string  `json:"models"`
		Transmissions []string  `json:"transmissions"`
		EngineVolumes []float64 `json:"engine_volumes"`
		FuelTypes     []string  `json:"fuel_types"`
		DriveTypes    []string  `json:"drive_types"`
	}
)
