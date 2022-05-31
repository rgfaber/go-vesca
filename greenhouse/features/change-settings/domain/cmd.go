package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core"
)

const CMD_TOPIC = "greenhouse:change-settings"

type Cmd struct {
	aggregateId core.IIdentity
	Settings    *model.Settings
}

func (c *Cmd) TraceId() string {
	//TODO implement me
	panic("implement me")
}

func (c *Cmd) AggregateId() core.IIdentity {
	return c.aggregateId
}

func NewCmd(aggregateId core.IIdentity, settings *model.Settings) *Cmd {
	return &Cmd{
		aggregateId: aggregateId,
		Settings:    settings,
	}
}
