package nats

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type NATSListener struct {
	natsBus INatsBus
	Handler interfaces.IFactHandler
	Topic   string
}

func (l *NATSListener) Activate() {
	l.natsBus.Listen(l.Topic, l.Handler)
}

func NewNATSListener(natsBus INatsBus,
	topic string,
	handler interfaces.IFactHandler) *NATSListener {
	return &NATSListener{
		natsBus: natsBus,
		Topic:   topic,
		Handler: handler,
	}
}
