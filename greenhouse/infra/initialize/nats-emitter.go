package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/sdk"
)

const FACT_TOPIC = "go-vesca.greenhouse.initialized"

type Emitter struct {
	natsBus infra.INatsBus
	memBus  sdk.IDECBus
}

func (e *Emitter) Emit(fact sdk.IFact) {
	e.natsBus.Publish(FACT_TOPIC, e.Serialize(fact))
}

func (e *Emitter) Bus() sdk.IDECBus {
	//TODO implement me
	panic("implement me")
}

func NewEmitter(memBus sdk.IDECBus, natsBus infra.INatsBus) *Emitter {
	return &Emitter{
		natsBus: natsBus,
		memBus:  memBus,
	}
}
