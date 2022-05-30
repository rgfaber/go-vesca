package contract

import "github.com/rgfaber/go-vesca/greenhouse/model"

type Hope struct {
	TargetId  string     `json:"target_id"`
	SessionId string     `json:"session_id"`
	Fan       *model.Fan `json:"fan"`
}

func (h *Hope) AggregateId() string {
	return h.TargetId
}

func (h *Hope) TraceId() string {
	return h.SessionId
}

func NewHope(id string, traceId string, fan *model.Fan, sprinkler *model.Sprinkler) *Hope {
	return &Hope{
		TargetId:  id,
		SessionId: traceId,
		Fan:       fan,
	}
}
