package bogus

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

var (
	MyFbkToJson = func(fbk interfaces.IFbk) ([]byte, error) {
		return json.Marshal(fbk)
	}
	MyFbkFromJson = func(data []byte) *MyFbk {
		var res *MyFbk
		json.Unmarshal(data, res)
		return res
	}
)

type MyFbk struct {
	error       string
	aggregateId string
	traceId     string
}

func (f *MyFbk) AggregateId() string {
	return f.aggregateId
}

func (f *MyFbk) TraceId() string {
	return f.traceId
}

func (f *MyFbk) IsSuccess() bool {
	return f.error != ""
}

func (f *MyFbk) Error() string {
	return f.error
}

func NewMyFbk(aggregateId string, traceId string, error string) *MyFbk {
	return &MyFbk{
		aggregateId: aggregateId,
		traceId:     traceId,
		error:       error,
	}
}
