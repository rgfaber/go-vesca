package initialize

import "github.com/rgfaber/go-vesca/th-sensor/model"

const CMD_TOPIC = "th_sensor:initialize"

type Cmd struct {
	measurement model.Measurement
	traceId     string
}

func NewCmd(traceId string, t, h float64) *Cmd {
	return &Cmd{
		traceId:     traceId,
		measurement: *model.NewMeasurement(t, h),
	}
}
