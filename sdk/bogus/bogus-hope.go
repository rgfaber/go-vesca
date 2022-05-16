package bogus

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type BogusHope struct {
	TargetId  string `json:"target_id"`
	SessionId string `json:"session_id"`
}

func (h *BogusHope) AggregateId() string {
	return h.TargetId
}

func (h *BogusHope) TraceId() string {
	return h.SessionId
}

var (
	BogusHopeFromJson = func(data []byte) interfaces.IHope {
		var res *BogusHope
		json.Unmarshal(data, res)
		return res
	}
	BogusHopeToJson = func(hope interfaces.IHope) ([]byte, error) {
		return json.Marshal(hope)
	}
)

func NewBogusHope(aggregateId string, traceId string) *BogusHope {
	return &BogusHope{
		TargetId:  aggregateId,
		SessionId: traceId,
	}
}
