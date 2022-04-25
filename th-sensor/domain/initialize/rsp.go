package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Rsp struct {
	aggregateId sdk.IIdentity
	status      model.SensorStatus
	isSuccess   bool
	traceId     string
}

func NewRsp(id sdk.IIdentity,
	status model.SensorStatus,
	isSuccess bool,
	traceId string) *Rsp {
	return &Rsp{
		aggregateId: id,
		status:      status,
		isSuccess:   isSuccess,
		traceId:     traceId,
	}
}

func (r *Rsp) AggregateId() sdk.IIdentity {
	return r.aggregateId
}

func (r *Rsp) Status() model.SensorStatus {
	return r.status
}

func (r *Rsp) IsSuccess() bool {
	return r.isSuccess
}

func (r *Rsp) TraceId() string {
	return r.traceId
}
