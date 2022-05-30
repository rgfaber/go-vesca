package bogus

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"log"
)

var MyFactFromJson = func(data []byte) domain.IFact {
	var res BogusFact
	json.Unmarshal(data, &res)
	return &res
}

var MyFactToJson = func(fct domain.IFact) []byte {
	//	f := fct.(*BogusFact)
	res, e := json.Marshal(fct)
	if e != nil {
		log.Fatal(e)
		return nil
	}
	return res
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
