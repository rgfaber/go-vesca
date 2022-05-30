package domain

import (
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"log"
)

type IEventHandler interface {
	Activate()
	Subscribe(events chan IEvt, transactional bool)
}

func ImplementsIEventHandler(handler IEventHandler) bool {
	return true
}

func BuildHandlers(handlers ...IEventHandler) []IEventHandler {
	return handlers
}

type EventHandler struct {
	MemBus mediator.IDECBus
	Topic  string
	Handle func(evt IEvt)
}

func (h *EventHandler) Activate() {
	h.MemBus.SubscribeAsync(h.Topic, h.Handle, false)
	h.MemBus.WaitAsync()
	log.Printf("Handling Mediator Event [%+v]\n", h.Topic)
}

func (h *EventHandler) Subscribe(events chan IEvt, transactional bool) {
	h.MemBus.SubscribeAsync(h.Topic, func(evt IEvt) {
		events <- evt
	}, transactional)
}

func NewEventHandler(memBus mediator.IDECBus, topic string, handle func(evt IEvt)) *EventHandler {
	return &EventHandler{
		MemBus: memBus,
		Topic:  topic,
		Handle: handle,
	}
}
