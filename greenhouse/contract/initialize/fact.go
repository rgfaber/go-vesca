package initialize

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
)

const FACT_TOPIC = "greenhouse.initialized"

type Fact struct {
	aggregateId string            `json:"aggregate_id"`
	traceId     string            `json:"trace_id"`
	payload     *model.Greenhouse `json:"payload"`
}

func (f *Fact) Serialize(fact sdk.IFact) string {
	s, err := json.Marshal(fact)
	if err != nil {
		panic(err)
	}
	return string(s)
}

func (f *Fact) Deserialize(s string) sdk.IFact {
	var ret Fact
	json.Unmarshal([]byte(s), &ret)
	return &ret
}

func (f *Fact) Topic() string {
	return FACT_TOPIC
}

func (f *Fact) AggregateId() string {
	return f.aggregateId
}

func (f *Fact) TraceId() string {
	return f.traceId
}

func (f *Fact) Payload() interface{} {
	return f.payload
}

func NewFact(aggregateId string, traceId string, greenhouse *model.Greenhouse) *Fact {
	return &Fact{
		aggregateId: aggregateId,
		traceId:     traceId,
		payload:     greenhouse,
	}
}
