package model

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/sdk/core"
)

type Sensor struct {
	ID      *core.Identity `json:"id"`
	Details *Details       `json:"details"`
}

func NewSensor(id string, name string) *Sensor {
	return &Sensor{
		ID:      core.NewIdentityFrom(config.GO_VESCA_TH_SENSOR_PREFIX, id),
		Details: NewDetails(name, ""),
	}
}
