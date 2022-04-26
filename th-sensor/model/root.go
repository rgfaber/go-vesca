package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	fan_status "github.com/rgfaber/go-vesca/th-sensor/model/fan-status"
	sprinkler_status "github.com/rgfaber/go-vesca/th-sensor/model/sprinkler-status"
)

type Root struct {
	SensorId        sdk.Identity                     `json:"id" bson:"id"`
	Greenhouse      Greenhouse                       `json:"greenhouse"`
	Details         Details                          `json:"details"`
	Measurement     Measurement                      `json:"measurement"`
	Status          SensorStatus                     `json:"status"`
	FanStatus       fan_status.FanStatus             `json:"fan_status"`
	SprinklerStatus sprinkler_status.SprinklerStatus `json:"sprinkler_status"`
}

func NewRoot(sensorId string, sensorName string, greenhouseId string) *Root {
	return &Root{
		SensorId:        *NewTHSensorId(sensorId),
		Greenhouse:      *NewGreenhouse(greenhouseId),
		Details:         *NewDetails(sensorName),
		Measurement:     *NewMeasurement(15, 23),
		Status:          Created,
		FanStatus:       fan_status.Deactivated,
		SprinklerStatus: sprinkler_status.Deactivated,
	}
}
