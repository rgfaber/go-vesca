package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
)

type Greenhouse struct {
	ID        *sdk.Identity    `json:"id" bson:"id"`
	Details   *Details         `json:"details"`
	Sensor    *Sensor          `json:"sensor"`
	Fan       *Fan             `json:"fan"`
	Sprinkler *Sprinkler       `json:"sprinkler"`
	Settings  *Settings        `json:"settings"`
	Status    GreenhouseStatus `json:"status"`
}

func NewGreenhouse(id string,
	name string,
	settings *Settings,
	sensor *Sensor,
	fan *Fan,
	sprinkler *Sprinkler) *Greenhouse {
	return &Greenhouse{
		ID:        sdk.NewIdentityFrom(config.GO_VESCA_GREENHOUSE_PREFIX, id),
		Details:   NewDetails(name, ""),
		Settings:  settings,
		Sensor:    sensor,
		Fan:       fan,
		Sprinkler: sprinkler,
		Status:    Created,
	}
}
