package domain

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
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
