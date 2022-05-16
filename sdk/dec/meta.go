package dec

import (
	"github.com/rgfaber/go-vesca/sdk/core"
)

type Meta struct {
	ID      core.Identity
	TraceId string
	Status  int
}

func NewMeta(id core.Identity, traceId string, status int) *Meta {
	return &Meta{
		ID:      id,
		TraceId: traceId,
		Status:  status,
	}
}
