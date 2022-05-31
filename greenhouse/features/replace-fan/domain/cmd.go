package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core"
)

type Cmd struct {
	ID      *core.Identity
	traceId string
	Fan     *model.Fan
}

func (c *Cmd) TraceId() string {
	return c.traceId
}

func NewCmd(id *core.Identity, traceId string,
	fan *model.Fan) *Cmd {
	return &Cmd{
		ID:      id,
		traceId: traceId,
		Fan:     fan,
	}
}

func (c *Cmd) AggregateId() core.IIdentity {
	return c.ID
}
