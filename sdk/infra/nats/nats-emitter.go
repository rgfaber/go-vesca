package nats

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type NatsEmitter struct {
	Handler interfaces.IEventHandler
	Topic   string
}

func (e *NatsEmitter) Activate() {
	e.Handler.Activate()
}

func NewNatsEmitter(natsBus INatsBus, topic string,
	handler interfaces.IEventHandler) *NatsEmitter {
	return &NatsEmitter{
		Handler: handler,
	}
}
