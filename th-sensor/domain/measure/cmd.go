package measure

import (
	"github.com/rgfaber/go-vesca/sdk"
)

const CMD_TOPIC = "th_sensor:measure"

type Cmd struct {
	aggregateId sdk.IIdentity
	traceId     string
}

func (c *Cmd) AggregateId() sdk.IIdentity {
	return c.aggregateId
}

func NewCmd(aggregateId sdk.IIdentity, traceId string) *Cmd {
	return &Cmd{
		aggregateId: aggregateId,
		traceId:     traceId,
	}
}
