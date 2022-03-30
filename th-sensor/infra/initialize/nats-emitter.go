package initialize

import (
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
)

type Emitter struct {
	events initialize.ChEvt
	bus    infra.IBus
}

func NewEmitter(bus infra.IBus, e initialize.ChEvt) *Emitter {
	return &Emitter{
		events: e,
		bus:    bus,
	}
}
