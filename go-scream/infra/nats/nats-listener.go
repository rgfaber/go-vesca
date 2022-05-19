package nats

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"sync"
)

type NATSListener struct {
	natsBus INatsBus
	Handler interfaces.IFactHandler
	Topic   string
	mutex   *sync.Mutex
}

func (l *NATSListener) Activate() {
	factChan := make(chan []byte)
	l.natsBus.Listen(l.Topic, factChan)
	go func(ch chan []byte) {
		for {
			data := <-ch
			l.mutex.Lock()
			l.Handler.Handle(data)
			l.mutex.Unlock()
		}
	}(factChan)
}

func NewNatsListener(natsBus INatsBus,
	topic string,
	handler interfaces.IFactHandler) *NATSListener {
	return &NATSListener{
		natsBus: natsBus,
		Topic:   topic,
		Handler: handler,
		mutex:   &sync.Mutex{},
	}
}
