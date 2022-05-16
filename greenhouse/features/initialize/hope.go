package initialize

import "github.com/rgfaber/go-vesca/greenhouse/model"

type Hope struct {
	TargetId  string           `json:"target_id"`
	SessionId string           `json:"session_id"`
	Fan       *model.Fan       `json:"fan"`
	Sprinkler *model.Sprinkler `json:"sprinkler"`
	Sensor    *model.Sensor    `json:"sensor"`
	Details   *model.Details   `json:"details"`
	Settings  *model.Settings  `json:"settings"`
}

func (h *Hope) AggregateId() string {
	return h.TargetId
}

func (h *Hope) TraceId() string {
	return h.SessionId
}

func NewHope(id string, traceId string, details *model.Details, settings *model.Settings, sensor *model.Sensor, fan *model.Fan, sprinkler *model.Sprinkler) *Hope {
	return &Hope{
		TargetId:  id,
		SessionId: traceId,
		Settings:  settings,
		Details:   details,
		Sensor:    sensor,
		Fan:       fan,
		Sprinkler: sprinkler,
	}
}
