package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
)

const CMD_TOPIC = "th_sensor:initialize"

type Cmd struct {
	ID        *sdk.Identity
	traceId   string
	Name      string
	Settings  *model.Settings
	Sensor    *model.Sensor
	Fan       *model.Fan
	Sprinkler *model.Sprinkler
}

func (c *Cmd) TraceId() string {
	return c.traceId
}

func NewCmd(id *sdk.Identity, traceId string, name string,
	settings *model.Settings,
	sensor *model.Sensor,
	fan *model.Fan,
	sprinkler *model.Sprinkler) *Cmd {
	return &Cmd{
		ID:        id,
		Name:      name,
		traceId:   traceId,
		Settings:  settings,
		Sensor:    sensor,
		Fan:       fan,
		Sprinkler: sprinkler,
	}
}

func (c *Cmd) AggregateId() sdk.IIdentity {
	return c.ID
}
