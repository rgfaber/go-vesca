package dec

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type Feature struct {
	Responders []interfaces.IHopeResponder
	Listeners  []interfaces.IFactListener
	Handlers   []interfaces.IEventHandler
}

func (f *Feature) Run() {
	for _, responder := range f.Responders {
		go responder.Activate()
	}
	for _, listener := range f.Listeners {
		go listener.Activate()
	}
	for _, handler := range f.Handlers {
		go handler.Activate()
	}
}

func NewFeature(
	responders []interfaces.IHopeResponder,
	listeners []interfaces.IFactListener,
	handlers []interfaces.IEventHandler) *Feature {
	return &Feature{
		Responders: responders,
		Listeners:  listeners,
		Handlers:   handlers,
	}
}
