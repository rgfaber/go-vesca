package bogus

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

var MyFactFromJson = func(data []byte) interfaces.IFact {
	var res BogusFact
	json.Unmarshal(data, &res)
	return &res
}

var MyFactToJson = func(fct interfaces.IFact) ([]byte, error) {
	//	f := fct.(*BogusFact)
	return json.Marshal(fct)
}

type BogusFact struct {
	OriginId  string `json:"origin_id"`
	SessionId string `json:"session_id"`
}

func NewBogusFact(aggregateId string, traceId string) *BogusFact {
	return &BogusFact{
		OriginId:  aggregateId,
		SessionId: traceId,
	}
}

func (f *BogusFact) AggregateId() string {
	return f.OriginId
}

func (f *BogusFact) TraceId() string {
	return f.SessionId
}
