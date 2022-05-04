package contract

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

const FACT_TOPIC = "greenhouse.initialized"

type Fact struct {
	AggregateId string            `json:"aggregate_id"`
	TraceId     string            `json:"trace_id"`
	Payload     *model.Greenhouse `json:"Payload"`
}

func NewFact(aggregateId string, traceId string, greenhouse *model.Greenhouse) *Fact {
	return &Fact{
		AggregateId: aggregateId,
		TraceId:     traceId,
		Payload:     greenhouse,
	}
}
