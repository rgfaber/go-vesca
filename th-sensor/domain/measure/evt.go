package measure

import (
	"github.com/rgfaber/go-vesca/sdk"
)

const EVT_TOPIC = "th_sensor:measured"

type Evt struct {
	aggregateId sdk.IIdentity
	temperature float64
	humidity    float64
}

func (e *Evt) AggregateId() sdk.IIdentity {
	return e.aggregateId

}

func NewEvt(aggregateId sdk.IIdentity, temperature float64, humidity float64) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		temperature: temperature,
		humidity:    humidity,
	}
}
