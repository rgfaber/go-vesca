package model

type Measurement struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func NewMeasurement(temperature float64, humidity float64) *Measurement {
	return &Measurement{
		Temperature: temperature,
		Humidity:    humidity,
	}
}
