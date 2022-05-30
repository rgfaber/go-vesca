package domain

import (
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
)

func PublishEvent(mediator mediator.IDECBus, topic string, evt domain.IEvt) {
	mediator.Publish(topic, evt)
}
