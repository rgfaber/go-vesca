package update_fan_status

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	fan_status "github.com/rgfaber/go-vesca/th-sensor/model/fan-status"
)

type Fbk struct {
	sensorId  string
	traceId   string
	isSuccess bool
	status    fan_status.FanStatus
}

func (f *Fbk) AggregateId() sdk.IIdentity {
	return model.NewTHSensorId(f.sensorId)
}

func (f *Fbk) Status() int {
	return int(f.status)
}

func (f *Fbk) IsSuccess() bool {
	return f.isSuccess
}

func (f *Fbk) TraceId() string {
	return f.traceId
}

func NewFbk(sensorId string, traceId string, isSuccess bool, status fan_status.FanStatus) *Fbk {
	return &Fbk{
		sensorId:  sensorId,
		traceId:   traceId,
		isSuccess: isSuccess,
		status:    status}
}
