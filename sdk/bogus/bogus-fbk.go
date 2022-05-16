package bogus

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

var (
	MyFbkToJson = func(fbk interfaces.IFbk) ([]byte, error) {
		return json.Marshal(fbk)
	}
	MyFbkFromJson = func(data []byte) *BogusFbk {
		var res *BogusFbk
		json.Unmarshal(data, res)
		return res
	}
)

type BogusFbk struct {
	ErrorMsg  string `json:"error_msg"`
	OriginId  string `json:"origin_id"`
	SessionId string `json:"session_id"`
}

func (f *BogusFbk) AggregateId() string {
	return f.OriginId
}

func (f *BogusFbk) TraceId() string {
	return f.SessionId
}

func (f *BogusFbk) IsSuccess() bool {
	return f.ErrorMsg != ""
}

func (f *BogusFbk) Error() string {
	return f.ErrorMsg
}

func NewBogusFbk(aggregateId string, traceId string, error string) *BogusFbk {
	return &BogusFbk{
		OriginId:  aggregateId,
		SessionId: traceId,
		ErrorMsg:  error,
	}
}
