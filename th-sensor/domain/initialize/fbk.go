package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Fbk struct {
	aggregateId sdk.IIdentity
	status      model.SensorStatus
	isSuccess   bool
	traceId     string
}

func NewFbk(id sdk.IIdentity,
	status model.SensorStatus,
	isSuccess bool,
	traceId string) *Fbk {
	return &Fbk{
		aggregateId: id,
		status:      status,
		isSuccess:   isSuccess,
		traceId:     traceId,
	}
}

func (r *Fbk) AggregateId() sdk.IIdentity {
	return r.aggregateId
}

func (r *Fbk) Status() int {
	return int(r.status)
}

func (r *Fbk) IsSuccess() bool {
	return r.isSuccess
}

func (r *Fbk) TraceId() string {
	return r.traceId
}
