package model

import "github.com/rgfaber/go-vesca/sdk"

type Root struct {
	SensorId    sdk.Identity `json:"id" bson:"id"`
	Greenhouse  Greenhouse   `json:"greenhouse"`
	Details     Details      `json:"details"`
	Measurement Measurement  `json:"measurement"`
	Status      SensorStatus `json:"status"`
}

func NewRoot(sensorId string, sensorName string, greenhouseId string) *Root {
	return &Root{
		SensorId:    *NewTHSensorId(sensorId),
		Greenhouse:  *NewGreenhouse(greenhouseId),
		Details:     *NewDetails(sensorName),
		Measurement: *NewMeasurement(15, 23),
		Status:      Created,
	}
}
