package nats

import (
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/sdk"
)

type Listener struct {
	memBus  sdk.IDECBus
	natsBus infra.INatsBus
}

func NewListener(memBus sdk.IDECBus, natsBus infra.INatsBus) Listener {
	return Listener{
		memBus:  memBus,
		natsBus: natsBus,
	}
}
