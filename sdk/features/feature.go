package features

import (
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra"
)

type IFeature interface {
	Run()
}

func ImplementsIFeature(feature IFeature) bool {
	return true
}

type Feature struct {
	Responders []infra.IHopeResponder
	Listeners  []infra.IFactListener
	Handlers   []domain.IEventHandler
}

func (f *Feature) Run() {
	for _, responder := range f.Responders {
		responder.Activate()
	}
	for _, listener := range f.Listeners {
		listener.Activate()
	}
	for _, handler := range f.Handlers {
		handler.Activate()
	}
}

func NewFeature(
	responders []infra.IHopeResponder,
	listeners []infra.IFactListener,
	handlers []domain.IEventHandler) *Feature {
	return &Feature{
		Responders: responders,
		Listeners:  listeners,
		Handlers:   handlers,
	}
}
