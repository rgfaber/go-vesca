package nats

import (
	"github.com/nats-io/nats.go"
	"github.com/rgfaber/go-vesca/sdk/infra"
	"log"
	"sync"
)

type NatsResponder struct {
	natsBus     INatsBus
	hopeHandler infra.IHopeHandler
	Topic       string
	mutex       *sync.Mutex
}

func (r *NatsResponder) Activate() {
	hopeChan := make(chan *nats.Msg)
	r.natsBus.Respond(r.Topic, hopeChan)
	go func(ch chan *nats.Msg) {
		for {
			r.mutex.Lock()
			msg := <-ch
			fbk := r.hopeHandler.Handle(msg.Data)
			err := msg.Respond(fbk)
			if err != nil {
				log.Fatal(err)
			}
			r.mutex.Unlock()
		}
	}(hopeChan)
}

func NewNatsResponder(natsBus INatsBus, topic string, handler infra.IHopeHandler) *NatsResponder {
	return &NatsResponder{
		natsBus:     natsBus,
		Topic:       topic,
		hopeHandler: handler,
		mutex:       &sync.Mutex{},
	}
}
