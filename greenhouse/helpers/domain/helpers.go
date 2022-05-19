package domain

import (
	interfaces2 "github.com/rgfaber/go-vesca/go-scream/infra/memory/interfaces"
	"github.com/rgfaber/go-vesca/go-scream/interfaces"
)

func PublishEvent(mediator interfaces2.IDECBus, topic string, evt interfaces.IEvt) {
	mediator.Publish(topic, evt)
}
