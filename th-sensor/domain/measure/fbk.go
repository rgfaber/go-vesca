package measure

import (
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Fbk struct {
	sensorId  string
	traceId   string
	isSuccess bool
	status    model.SensorStatus
}

func (f *Fbk) AggregateId() string {
	return f.sensorId
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

func NewFbk(sensorId string, traceId string, isSuccess bool, status model.SensorStatus) *Fbk {
	return &Fbk{
		sensorId:  sensorId,
		traceId:   traceId,
		isSuccess: isSuccess,
		status:    status,
	}
}
