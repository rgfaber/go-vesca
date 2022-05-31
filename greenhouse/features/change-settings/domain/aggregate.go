package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/helpers/infra"
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
	"log"
)

type Aggregate struct {
	Store  store.IStore
	MemBus mediator.IDECBus
}

func (a Aggregate) LoadEvents(aggregateId core.IIdentity) []domain.IEvt {
	//TODO implement me
	panic("implement me")
}

func (a Aggregate) AppendEvent(evt domain.IEvt) {
	//TODO implement me
	panic("implement me")
}

func (a Aggregate) Attempt(cmd domain.ICmd) (domain.IFbk, error) {
	if cmd == nil {
		return nil, nil
	}
	//	c := cmd.(*Cmd)
	state := infra.LoadGreenhouse(a.Store, cmd.AggregateId().Id())
	log.Println(state)
	return nil, nil

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
