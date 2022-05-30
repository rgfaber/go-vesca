package domain

import (
	"github.com/rgfaber/go-vesca/sdk/core"
)

const CMD_TOPIC = "th_sensor:measure-th"

type Cmd struct {
	aggregateId core.Identity
	traceId     string
}

func (c *Cmd) TraceId() string {
	return c.traceId
}

func (c *Cmd) AggregateId() core.IIdentity {
	return &c.aggregateId
}

func NewCmd(aggregateId core.Identity, traceId string) *Cmd {
	return &Cmd{
		aggregateId: aggregateId,
		traceId:     traceId,
	}
}
