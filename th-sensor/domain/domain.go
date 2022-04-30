package domain

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

func LoadState(store dec.IStore, aggregateId string) *model.Root {
	return store.Load(aggregateId).(*model.Root)
}

func SaveState(store dec.IStore, state *model.Root) {
	store.Save(state)
}

func Publish(bus dec.IDECBus, topic string, evt dec.IEvt) {
	bus.Publish(topic, evt)
}
