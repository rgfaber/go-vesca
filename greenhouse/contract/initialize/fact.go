package initialize

import "github.com/rgfaber/go-vesca/th-sensor/model"

const FACT_TOPIC = "greenhouse.initialized"

type Fact struct {
	aggregateId string            `json:"aggregate_id"`
	traceId     string            `json:"trace_id"`
	payload     *model.Greenhouse `json:"payload"`
}

func (f Fact) Topic() string {
	return FACT_TOPIC
}

func (f Fact) AggregateId() string {
	return f.aggregateId
}

func (f Fact) TraceId() string {
	return f.traceId
}

func (f Fact) Payload() interface{} {
	return f.payload
}

func NewFact(aggregateId string, traceId string, greenhouse *model.Greenhouse) Fact {
	return Fact{
		aggregateId: aggregateId,
		traceId:     traceId,
		payload:     greenhouse,
	}
}
