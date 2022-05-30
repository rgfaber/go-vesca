package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/helpers/infra"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
)

type Aggregate struct {
	Store  store.IStore
	MemBus mediator.IDECBus
}

func (a Aggregate) Attempt(cmd domain.ICmd) (domain.IFbk, error) {
	if cmd == nil {
		return nil, nil
	}
	c := cmd.(*Cmd)
	state := infra.LoadGreenhouse(a.Store, cmd.AggregateId().Id())

}

func (a Aggregate) Apply(evt domain.IEvt) {
	//TODO implement me
	panic("implement me")
}

func NewAggregate(memBus mediator.IDECBus, store store.IStore) *Aggregate {
	return &Aggregate{
		Store:  store,
		MemBus: memBus,
	}
}
