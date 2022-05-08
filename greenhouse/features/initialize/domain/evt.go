package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
)

const EVT_TOPIC = "greenhouse:initialized"

type Evt struct {
	aggregateId sdk.IIdentity
	traceId     string
	settings    *model.Settings
}

func (e *Evt) TraceId() string {
	return e.traceId
}

func NewEvt(aggregateId sdk.IIdentity, traceId string, settings *model.Settings) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		traceId:     traceId,
		settings:    settings,
	}
}

func (e *Evt) AggregateId() sdk.IIdentity {
	return e.aggregateId
}
