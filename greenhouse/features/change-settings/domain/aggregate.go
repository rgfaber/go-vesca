package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/sdk"
)

type Aggregate struct {
	Store  sdk.IStore
	MemBus sdk.IDECBus
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

func NewAggregate(memBus sdk.IDECBus, store sdk.IStore) *Aggregate {
	return &Aggregate{
		Store:  store,
		MemBus: memBus,
	}
}
