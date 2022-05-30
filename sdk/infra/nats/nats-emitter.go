package nats

import (
	"github.com/rgfaber/go-vesca/sdk/domain"
)

type NatsEmitter struct {
	natsBus INatsBus
	Handler domain.IEventHandler
	Topic   string
	//mutex   *sync.Mutex
	//evt2Fact func(evt interfaces.IEvt) interfaces.IFact
	//fact2Str func(fact interfaces.IFact) []byte
}

func (e *NatsEmitter) Activate() {
	e.Handler.Activate()
	//	events := make(chan interfaces.IEvt)
	//	e.Handler.Subscribe(events)
	//	go func(ch chan interfaces.IEvt) {
	//		for {
	//			evt := <-events
	//			e.mutex.Lock()
	//			fact := e.evt2Fact(evt)
	//			data := e.fact2Str(fact)
	//			e.natsBus.Publish(e.Topic, data)
	//			e.mutex.Unlock()
	//		}
	//	}(events)
	////	select {}
	//	//e.Handler.Activate(events)
}

func NewNatsEmitter(natsBus INatsBus, topic string,
	handler domain.IEventHandler) *NatsEmitter {
	return &NatsEmitter{
		natsBus: natsBus,
		Topic:   topic,
		Handler: handler,
		//evt2Fact: evt2Fact,
		//fact2Str: fact2Str,
		//mutex: &sync.Mutex{},
	}
}
