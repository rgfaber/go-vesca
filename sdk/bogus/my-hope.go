package bogus

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type MyHope struct {
	aggregateId string
	traceId     string
}

func (h *MyHope) AggregateId() string {
	return h.aggregateId
}

func (h *MyHope) TraceId() string {
	return h.traceId
}

var (
	MyHopeFromJson = func(data []byte) interfaces.IHope {
		var res *MyHope
		json.Unmarshal(data, res)
		return res
	}
	MyHopeToJson = func(hope interfaces.IHope) ([]byte, error) {
		return json.Marshal(hope)
	}
)

func NewMyHope(aggregateId string, traceId string) *MyHope {
	return &MyHope{
		aggregateId: aggregateId,
		traceId:     traceId,
	}
}
