package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core"
)

type Evt struct {
	aggregateId core.IIdentity
	traceId     string
	Fan         *model.Fan
}

func (e *Evt) TraceId() string {
	return e.traceId
}

func NewEvt(aggregateId core.IIdentity, traceId string, fan *model.Fan) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		traceId:     traceId,
		Fan:         fan,
	}
}

func (e *Evt) AggregateId() core.IIdentity {
	return e.aggregateId
}
