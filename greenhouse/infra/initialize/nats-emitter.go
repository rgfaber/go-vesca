package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/sdk/dec"
)

type Emitter struct {
	bus infra.INatsBus
}

func (e *Emitter) Emit(fact dec.IFact) {
	e.bus.Publish(fact.Topic(), fact.AsJson())
}

func (e *Emitter) Bus() dec.IDECBus {
	//TODO implement me
	panic("implement me")
}

func NewEmitter(bus infra.INatsBus) *Emitter {
	return &Emitter{
		bus: bus,
	}
}
