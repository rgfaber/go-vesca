package bogus

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

var MyFactFromJson = func(data []byte) interfaces.IFact {
	var res MyFact
	json.Unmarshal(data, &res)
	return &res
}

var MyFactToJson = func(fct interfaces.IFact) ([]byte, error) {
	//	f := fct.(*MyFact)
	return json.Marshal(fct)
}

type MyFact struct {
	OriginId  string `json:"origin_id"`
	SessionId string `json:"session_id"`
}

func NewMyFact(aggregateId string, traceId string) *MyFact {
	return &MyFact{
		OriginId:  aggregateId,
		SessionId: traceId,
	}
}

func (f *MyFact) AggregateId() string {
	return f.OriginId
}

func (f *MyFact) TraceId() string {
	return f.SessionId
}
