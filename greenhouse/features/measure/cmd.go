package measure

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

const CMD_TOPIC = "th_sensor:measure"

type Cmd struct {
	aggregateId core.Identity
	traceId     string
}

func (c *Cmd) TraceId() string {
	return c.traceId
}

func (c *Cmd) AggregateId() interfaces.IIdentity {
	return &c.aggregateId
}

func NewCmd(aggregateId core.Identity, traceId string) *Cmd {
	return &Cmd{
		aggregateId: aggregateId,
		traceId:     traceId,
	}
}
