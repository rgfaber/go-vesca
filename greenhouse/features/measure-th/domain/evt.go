package domain

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

const EVT_TOPIC = "th_sensor:measured"

type Evt struct {
	aggregateId interfaces.IIdentity
	traceId     string
	temperature float64
	humidity    float64
}

func (e *Evt) TraceId() string {
	return e.traceId
}

func (e *Evt) AggregateId() interfaces.IIdentity {
	return e.aggregateId
}

func NewEvt(aggregateId interfaces.IIdentity, traceId string, temperature float64, humidity float64) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		traceId:     traceId,
		temperature: temperature,
		humidity:    humidity,
	}
}
