package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

type Fbk struct {
	aggregateId string
	status      model.GreenhouseStatus
	traceId     string
	error       string
}

func (f *Fbk) Error() string {
	return f.error
}

func NewFbk(id string,
	status model.GreenhouseStatus,
	traceId string,
	error string) *Fbk {
	return &Fbk{
		aggregateId: id,
		status:      status,
		traceId:     traceId,
		error:       error,
	}
}

func (f *Fbk) AggregateId() string {
	return f.aggregateId
}

func (f *Fbk) Status() int {
	return int(f.status)
}

func (f *Fbk) IsSuccess() bool {
	return f.error != ""
}

func (f *Fbk) TraceId() string {
	return f.traceId
}
