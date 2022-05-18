package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

type Fbk struct {
	Aggregate_Id      string                 `json:"aggregate_id"`
	Greenhouse_Status model.GreenhouseStatus `json:"greenhouse_status"`
	Trace_Id          string                 `json:"trace_id"`
	Err               string                 `json:"err"`
}

func (f *Fbk) Error() string {
	return f.Err
}

func NewFbk(id string,
	status model.GreenhouseStatus,
	traceId string,
	error string) *Fbk {
	return &Fbk{
		Aggregate_Id:      id,
		Greenhouse_Status: status,
		Trace_Id:          traceId,
		Err:               error,
	}
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
