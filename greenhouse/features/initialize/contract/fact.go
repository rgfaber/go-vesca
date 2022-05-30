package contract

import (
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/topics"
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

type Fact struct {
	Aggregate_Id string            `json:"aggregate_id"`
	Trace_Id     string            `json:"trace_id"`
	Payload      *model.Greenhouse `json:"payload"`
}

func (f *Fact) AggregateId() string {
	return f.Aggregate_Id
}

func (f *Fact) TraceId() string {
	return f.Trace_Id
}

func (f *Fact) Topic() string {
	return topics.NATS_FACT_TOPIC
}

func NewFact(aggregateId string, traceId string, greenhouse *model.Greenhouse) *Fact {
	return &Fact{
		Aggregate_Id: aggregateId,
		Trace_Id:     traceId,
		Payload:      greenhouse,
	}
}
