package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/dec"
)

func LoadState(store dec.IStore, aggregateId string) *model.Greenhouse {
	return store.Load(aggregateId).(*model.Greenhouse)
}

func SaveState(store dec.IStore, state *model.Greenhouse) {
	store.Save(state)
}

func Publish(bus dec.IDECBus, topic string, evt dec.IEvt) {
	bus.Publish(topic, evt)
}
