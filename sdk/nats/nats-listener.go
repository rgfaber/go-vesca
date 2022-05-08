package nats

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
)

type Listener struct {
	natsBus INatsBus
	handler dec.IFactHandler
	Topic   string
}

func (l *Listener) Listen() {
	l.natsBus.Listen(l.Topic, l.handler)
}

func NewListener(natsBus INatsBus,
	topic string,
	handler dec.IFactHandler) *Listener {
	return &Listener{
		natsBus: natsBus,
		Topic:   topic,
		handler: handler,
	}
}
