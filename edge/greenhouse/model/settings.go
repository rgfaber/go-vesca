package model

type Settings struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func NewSettings(temperature float64, humidity float64) *Settings {
	return &Settings{
		Temperature: temperature,
		Humidity:    humidity,
	}
}
