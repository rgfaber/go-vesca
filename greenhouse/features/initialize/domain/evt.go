package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core"
)

type Evt struct {
	aggregateId core.IIdentity
	traceId     string
	Greenhouse  *model.Greenhouse
}

func (e *Evt) TraceId() string {
	return e.traceId
}

func NewEvt(aggregateId core.IIdentity, traceId string, greenhouse *model.Greenhouse) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		traceId:     traceId,
		Greenhouse:  greenhouse,
	}
}

func (e *Evt) AggregateId() core.IIdentity {
	return e.aggregateId
}
