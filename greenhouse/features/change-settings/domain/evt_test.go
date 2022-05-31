package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core"
)

const EVT_TOPIC = "greenhouse:settings-changed"

type Evt struct {
	aggregateId core.IIdentity
	Settings    *model.Settings
}

func NewEvt(aggregateId core.IIdentity, settings *model.Settings) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		Settings:    settings,
	}
}
