package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

type Fbk struct {
	Aggregate_Id      string                 `json:"aggregate_id"`
	Trace_Id          string                 `json:"trace_id"`
	Err               string                 `json:"err"`
	Greenhouse_Status model.GreenhouseStatus `json:"greenhouse_status"`
}

func (f *Fbk) Error() string {
	return f.Err
}

func (f *Fbk) AggregateId() string {
	return f.Aggregate_Id
}

func (f *Fbk) Status() int {
	return int(f.Greenhouse_Status)
}

func (f *Fbk) IsSuccess() bool {
	return f.Err == ""
}

func (f *Fbk) TraceId() string {
	return f.Trace_Id
}

func NewFbk(sensorId string, traceId string, status model.GreenhouseStatus, err string) *Fbk {
	return &Fbk{
		Aggregate_Id:      sensorId,
		Trace_Id:          traceId,
		Greenhouse_Status: status,
		Err:               err,
	}
}
