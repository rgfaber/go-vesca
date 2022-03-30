package initialize

import (
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
)

type Emitter struct {
	events initialize.Events
	bus    infra.IBus
}

func NewEmitter(bus infra.IBus, e initialize.Events) *Emitter {
	return &Emitter{
		events: e,
		bus:    bus,
	}
}
