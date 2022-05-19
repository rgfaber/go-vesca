package domain

import (
	"github.com/rgfaber/go-vesca/go-scream/core"
	"github.com/rgfaber/go-vesca/go-scream/interfaces"
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

type Cmd struct {
	ID        *core.Identity
	traceId   string
	Details   *model.Details
	Settings  *model.Settings
	Sensor    *model.Sensor
	Fan       *model.Fan
	Sprinkler *model.Sprinkler
}

func (c *Cmd) TraceId() string {
	return c.traceId
}

func NewCmd(id *core.Identity, traceId string,
	details *model.Details,
	settings *model.Settings,
	sensor *model.Sensor,
	fan *model.Fan,
	sprinkler *model.Sprinkler) *Cmd {
	return &Cmd{
		ID:        id,
		Details:   details,
		traceId:   traceId,
		Settings:  settings,
		Sensor:    sensor,
		Fan:       fan,
		Sprinkler: sprinkler,
	}
}

func (c *Cmd) AggregateId() interfaces.IIdentity {
	return c.ID
}
