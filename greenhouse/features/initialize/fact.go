package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

type Fact struct {
	aggregateId string            `json:"aggregate_id"`
	traceId     string            `json:"trace_id"`
	Payload     *model.Greenhouse `json:"Payload"`
}

func (f Fact) AggregateId() string {
	return f.aggregateId
}

func (f Fact) TraceId() string {
	return f.traceId
}

func (f Fact) Topic() string {
	return NATS_FACT_TOPIC
}

func NewFact(aggregateId string, traceId string, greenhouse *model.Greenhouse) *Fact {
	return &Fact{
		aggregateId: aggregateId,
		traceId:     traceId,
		Payload:     greenhouse,
	}
}
