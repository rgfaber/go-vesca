package nats

import (
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/greenhouse/infra/initialize/contract"
	"github.com/rgfaber/go-vesca/sdk"
)

const FACT_TOPIC = "go-vesca.greenhouse.initialized"

type Emitter struct {
	natsBus    infra.INatsBus
	memBus     sdk.IDECBus
	serializer Serializer
}

func (e *Emitter) Topic() string {
	return FACT_TOPIC
}

func (e *Emitter) Emit(fact sdk.IFact) {
	f := fact.(*contract.Fact)
	e.natsBus.Publish(e.Topic(), e.serializer.Serialize(f))
}

func NewEmitter(memBus sdk.IDECBus, natsBus infra.INatsBus) Emitter {
	return Emitter{
		natsBus:    natsBus,
		memBus:     memBus,
		serializer: NewSerializer(),
	}
}
