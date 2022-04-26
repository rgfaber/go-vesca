package initialize

import (
	"github.com/rgfaber/go-vesca/th-sensor/infra"
)

type IEmitter interface {
	Emit(fact IFact)
}

type Emitter struct {
	bus infra.IBus
}

func NewEmitter(bus infra.IBus) *Emitter {
	return &Emitter{
		bus: bus,
	}
}
