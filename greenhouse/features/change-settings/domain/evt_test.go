package domain

const EVT_TOPIC = "greenhouse:settings-changed"

type Evt struct {
	aggregateId sdk.IIdentity
	Settings    *model.Settings
}

func NewEvt(aggregateId sdk.IIdentity, settings *model.Settings) *Evt {
	return &Evt{
		aggregateId: aggregateId,
		Settings:    settings,
	}
}
