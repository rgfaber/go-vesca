package nats

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/nats"
)

const FACT_TOPIC = "govesca.greenhouse.measured"

type Listener struct {
	memBus     sdk.IDECBus
	natsBus    nats.INatsBus
	serializer sdk.IFactSerializer
}

func (l Listener) Listen(f func(fact sdk.IFact)) {
	l.natsBus.Listen(FACT_TOPIC, func(data []byte) {
		f := l.serializer.Deserialize(data)
		fmt.Println(f)
	})
}

func NewListener(memBus sdk.IDECBus, natsBus nats.INatsBus) Listener {
	return Listener{
		memBus:  memBus,
		natsBus: natsBus,
	}
}
