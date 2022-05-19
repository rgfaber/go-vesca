package dec

import (
	"github.com/rgfaber/go-vesca/sdk/infra/memory/interfaces"
	sdk_interfaces "github.com/rgfaber/go-vesca/sdk/interfaces"
)

type EventHandler struct {
	MemBus interfaces.IDECBus
	Topic  string
	Handle func(evt sdk_interfaces.IEvt)
}

func (h *EventHandler) Activate() {
	h.MemBus.SubscribeAsync(h.Topic, h.Handle, false)
	h.MemBus.WaitAsync()
}

func (h *EventHandler) Subscribe(events chan sdk_interfaces.IEvt, transactional bool) {
	h.MemBus.SubscribeAsync(h.Topic, func(evt sdk_interfaces.IEvt) {
		events <- evt
	}, transactional)
}

func NewEventHandler(memBus interfaces.IDECBus, topic string, handle func(evt sdk_interfaces.IEvt)) *EventHandler {
	return &EventHandler{
		MemBus: memBus,
		Topic:  topic,
		Handle: handle,
	}
}
