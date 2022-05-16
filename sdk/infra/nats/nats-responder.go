package nats

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type NatsResponder struct {
	natsBus     INatsBus
	hopeHandler interfaces.IHopeHandler
	Topic       string
}

func (r *NatsResponder) Activate() {
	r.natsBus.Respond(r.Topic, r.hopeHandler)
}

func NewNatsResponder(natsBus INatsBus, topic string, handler interfaces.IHopeHandler) *NatsResponder {
	return &NatsResponder{
		natsBus:     natsBus,
		Topic:       topic,
		hopeHandler: handler,
	}
}
