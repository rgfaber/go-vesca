package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Rsp struct {
	AggregateId sdk.Identity
	Status      model.SensorStatus
	IsSuccess   bool
	TraceId     string
}

func NewRsp(id sdk.Identity,
	status model.SensorStatus,
	isSucces bool,
	traceId string) *Rsp {
	return &Rsp{
		AggregateId: id,
		Status:      status,
		IsSuccess:   isSucces,
		TraceId:     traceId,
	}
}
