package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

const CMD_TOPIC = "greenhouse:change-settings"

type Cmd struct {
	aggregateId sdk.IIdentity
	Settings    *model.Settings
}

func (c *Cmd) AggregateId() sdk.IIdentity {
	return c.aggregateId
}

func NewCmd(aggregateId sdk.IIdentity, settings *model.Settings) *Cmd {
	return &Cmd{
		aggregateId: aggregateId,
		Settings:    settings,
	}
}
