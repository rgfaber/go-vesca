package nats

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
)

type NatsResponder struct {
	natsBus     INatsBus
	hopeHandler dec.IHopeHandler
}

func NewResponder(natsBus INatsBus,
	hopeHandler dec.IHopeHandler,
) *NatsResponder {
	return &NatsResponder{
		natsBus:     natsBus,
		hopeHandler: hopeHandler,
	}
}

func (r *NatsResponder) Respond(topic string) {
	r.natsBus.Respond(topic, r.hopeHandler)
}
