package domain

import (
	infra2 "github.com/rgfaber/go-vesca/go-scream/infra"
	"github.com/rgfaber/go-vesca/greenhouse/infra"
)

type Aggregate struct {
	Store  infra2.IStore
	MemBus infra2.IDECBus
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

func NewAggregate(memBus infra2.IDECBus, store infra2.IStore) *Aggregate {
	return &Aggregate{
		Store:  store,
		MemBus: memBus,
	}
}
