package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

const EVT_TOPIC = "th_sensor:initialized"

type Evt struct {
	aggregateId sdk.Identity
	traceId     string
	measurement model.Measurement
}

func NewEvt(aggregateId sdk.Identity, traceId string, measurement model.Measurement) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		traceId:     traceId,
		measurement: measurement,
	}
}
