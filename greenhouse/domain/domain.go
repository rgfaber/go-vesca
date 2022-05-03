package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
)

func LoadState(store sdk.IStore, aggregateId string) *model.Greenhouse {
	return store.Load(aggregateId).(*model.Greenhouse)
}

func SaveState(store sdk.IStore, state *model.Greenhouse) {
	store.Save(state)
}

func Publish(bus sdk.IDECBus, topic string, evt sdk.IEvt) {
	bus.Publish(topic, evt)
}
