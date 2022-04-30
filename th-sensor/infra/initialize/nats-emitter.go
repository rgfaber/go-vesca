package initialize

import (
	"github.com/rgfaber/go-vesca/th-sensor/infra"
)

type IEmitter interface {
	Emit(fact IFact)
}

type Emitter struct {
	bus infra.INatsBus
}

func NewEmitter(bus infra.INatsBus) *Emitter {
	return &Emitter{
		bus: bus,
	}
}
