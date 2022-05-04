package nats

import (
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/greenhouse/infra/initialize/contract"
	"github.com/rgfaber/go-vesca/sdk"
)

const FACT_TOPIC = "go-vesca.greenhouse.measured"

type Emitter struct {
	natsBus    infra.INatsBus
	memBus     sdk.IDECBus
	Serializer Serializer
}

func (e *Emitter) Serialize(fact sdk.IFact) []byte {
	//TODO implement me
	f := fact.(*contract.Fact)
	return e.Serialize(f)
}

func (e *Emitter) Deserialize(s []byte) sdk.IFact {
	//TODO implement me
	panic("implement me")
}

func (e *Emitter) Topic() string {
	return FACT_TOPIC
}

func (e *Emitter) Emit(fact sdk.IFact) {
	f := fact.(*contract.Fact)
	e.natsBus.Publish(e.Topic(), e.Serializer.Serialize(f))
}

func (e *Emitter) Bus() sdk.IDECBus {
	//TODO implement me
	panic("implement me")
}

func NewEmitter(memBus sdk.IDECBus, natsBus infra.INatsBus) Emitter {
	return Emitter{
		natsBus:    natsBus,
		memBus:     memBus,
		Serializer: NewSerializer(),
	}
}
