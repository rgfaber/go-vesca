package nats

import (
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/contract"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/nats"
)

type Emitter struct {
	natsBus    nats.INatsBus
	memBus     sdk.IDECBus
	serializer sdk.IFactSerializer
	Topic      string
}

func (e *Emitter) Emit(fact sdk.IFact) {
	f := fact.(*contract.Fact)
	s, err := e.serializer.Serialize(f)
	if err != nil {
		panic(err)
	}
	e.natsBus.Publish(e.Topic, s)
}

func NewEmitter(memBus sdk.IDECBus, natsBus nats.INatsBus, topic string, serializer sdk.IFactSerializer) *Emitter {
	return &Emitter{
		natsBus:    natsBus,
		memBus:     memBus,
		serializer: serializer,
		Topic:      topic,
	}
}
