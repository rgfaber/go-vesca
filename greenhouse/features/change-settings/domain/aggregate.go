package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/helpers/infra"
	sdk_infra "github.com/rgfaber/go-vesca/sdk/infra"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/interfaces"
)

type Aggregate struct {
	Store  interfaces.IStore
	MemBus sdk_infra.IDECBus
}

func (a Aggregate) Attempt(cmd sdk.ICmd) (sdk.IFbk, error) {
	if cmd == nil {
		return nil, nil
	}
	c := cmd.(*Cmd)
	state := infra.LoadGreenhouse(a.Store, cmd.AggregateId().Id())

}

func (a Aggregate) Apply(evt sdk.IEvt) {
	//TODO implement me
	panic("implement me")
}

func NewAggregate(memBus sdk_infra.IDECBus, store sdk_infra.IStore) *Aggregate {
	return &Aggregate{
		Store:  store,
		MemBus: memBus,
	}
}
