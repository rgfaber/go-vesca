package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
)

const EVT_TOPIC = "th_sensor:initialized"

type Evt struct {
	aggregateId sdk.IIdentity
	traceId     string
	measurement model.Settings
}

func NewEvt(aggregateId sdk.IIdentity, traceId string, measurement model.Settings) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		traceId:     traceId,
		measurement: measurement,
	}
}

func (e *Evt) AggregateId() sdk.IIdentity {
	return e.aggregateId
}
