package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type Evt struct {
	aggregateId interfaces.IIdentity
	traceId     string
	Greenhouse  *model.Greenhouse
}

func (e *Evt) TraceId() string {
	return e.traceId
}

func NewEvt(aggregateId interfaces.IIdentity, traceId string, greenhouse *model.Greenhouse) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		traceId:     traceId,
		Greenhouse:  greenhouse,
	}
}

func (e *Evt) AggregateId() interfaces.IIdentity {
	return e.aggregateId
}
