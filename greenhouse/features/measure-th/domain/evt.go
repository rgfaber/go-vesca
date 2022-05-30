package domain

import (
	"github.com/rgfaber/go-vesca/sdk/core"
)

const EVT_TOPIC = "th_sensor:measured"

type Evt struct {
	aggregateId core.IIdentity
	traceId     string
	temperature float64
	humidity    float64
}

func (e *Evt) Order() int {
	//TODO implement me
	panic("implement me")
}

func (e *Evt) TraceId() string {
	return e.traceId
}

func (e *Evt) AggregateId() core.IIdentity {
	return e.aggregateId
}

func NewEvt(aggregateId core.IIdentity, traceId string, temperature float64, humidity float64) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		traceId:     traceId,
		temperature: temperature,
		humidity:    humidity,
	}
}
