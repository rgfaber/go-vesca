package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

const CMD_TOPIC = "th_sensor:initialize"

type Cmd struct {
	aggregateId sdk.Identity
	measurement model.Measurement
	traceId     string
}

func NewCmd(aggregateId sdk.Identity, traceId string, t, h float64) *Cmd {
	return &Cmd{
		aggregateId: aggregateId,
		traceId:     traceId,
		measurement: *model.NewMeasurement(t, h),
	}
}
