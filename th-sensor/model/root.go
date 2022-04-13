package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
)

type Root struct {
	ID          sdk.Identity `json:"id" bson:"id"`
	Greenhouse  Greenhouse   `json:"greenhouse"`
	Details     Details      `json:"details"`
	Measurement Measurement  `json:"measurement"`
	Status      SensorStatus `json:"status"`
}

func NewRoot(cfg *config.Config) *Root {
	return &Root{
		ID:          *NewTHSensorId(cfg),
		Greenhouse:  *NewGreenhouse(cfg.GreenhouseId()),
		Details:     *NewDetails(cfg.SensorName()),
		Measurement: *NewMeasurement(15, 23),
		Status:      Created,
	}
}
