package model

import (
	"github.com/rgfaber/go-vesca/go-scream/core"
	"github.com/rgfaber/go-vesca/greenhouse/config"
)

type Greenhouse struct {
	ID        *core.Identity   `json:"id" bson:"id"`
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
		ID:        core.NewIdentityFrom(config.GO_VESCA_GREENHOUSE_PREFIX, id),
		Details:   NewDetails(name, ""),
		Settings:  settings,
		Sensor:    sensor,
		Fan:       fan,
		Sprinkler: sprinkler,
		Status:    Created,
	}
}
