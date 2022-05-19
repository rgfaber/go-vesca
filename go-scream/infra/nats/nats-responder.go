package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"log"
	"sync"
)

type NatsResponder struct {
	natsBus     INatsBus
	hopeHandler interfaces.IHopeHandler
	Topic       string
	mutex       *sync.Mutex
}

func (r *NatsResponder) Activate() {
	hopeChan := make(chan *nats.Msg)
	r.natsBus.Respond(r.Topic, hopeChan)
	go func(ch chan *nats.Msg) {
		for {
			msg := <-ch
			r.mutex.Lock()
			fbk := r.hopeHandler.Handle(msg.Data)
			err := msg.Respond(fbk)
			if err != nil {
				log.Fatal(err)
			}
			r.mutex.Unlock()
		}
	}(hopeChan)
}

func NewNatsResponder(natsBus INatsBus, topic string, handler interfaces.IHopeHandler) *NatsResponder {
	return &NatsResponder{
		natsBus:     natsBus,
		Topic:       topic,
		hopeHandler: handler,
		mutex:       &sync.Mutex{},
	}
}
