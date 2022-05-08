package domain

import (
	"github.com/rgfaber/go-vesca/sdk"
)

const CMD_TOPIC = "th_sensor:measure"

type Cmd struct {
	aggregateId sdk.Identity
	traceId     string
}

func (c *Cmd) AggregateId() sdk.IIdentity {
	return &c.aggregateId
}

func NewCmd(aggregateId sdk.Identity, traceId string) *Cmd {
	return &Cmd{
		aggregateId: aggregateId,
		traceId:     traceId,
	}
}
