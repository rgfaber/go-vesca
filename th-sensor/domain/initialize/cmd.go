package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

const CMD_TOPIC = "th_sensor:initialize"

type Cmd struct {
	measurement  model.Measurement
	traceId      string
	SensorId     string
	SensorName   string
	GreenhouseId string
}

func NewCmd(sensorId string, sensorName string, greenhouseId string, traceId string, t, h float64) *Cmd {
	return &Cmd{
		SensorId:     sensorId,
		SensorName:   sensorName,
		GreenhouseId: greenhouseId,
		traceId:      traceId,
		measurement:  *model.NewMeasurement(t, h),
	}
}

func (c *Cmd) AggregateId() sdk.IIdentity {
	return model.NewTHSensorId(c.SensorId)
}
