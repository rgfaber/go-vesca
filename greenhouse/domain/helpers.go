package domain

import "github.com/rgfaber/go-vesca/sdk"

func PublishEvent(bus sdk.IDECBus, topic string, evt sdk.IEvt) {
	bus.Publish(topic, evt)
}
