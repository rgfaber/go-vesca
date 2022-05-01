package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

const CMD_TOPIC = "th_sensor:initialize"

type Cmd struct {
	measurement  model.Measurement
	aggregateId  sdk.Identity
	traceId      string
	sensorName   string
	greenhouseId string
}

func NewCmd(aggregateId sdk.Identity, sensorName string, greenhouseId string, traceId string, t, h float64) *Cmd {
	return &Cmd{
		aggregateId:  aggregateId,
		sensorName:   sensorName,
		greenhouseId: greenhouseId,
		traceId:      traceId,
		measurement:  *model.NewMeasurement(t, h),
	}
}

func (c *Cmd) AggregateId() sdk.IIdentity {
	return &c.aggregateId
}
