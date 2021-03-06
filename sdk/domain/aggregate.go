package domain

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
)

type AggregateBase struct {
	EventStream IEventStream
	Mediator    mediator.IDECBus
	State       interface{}
}

func (a *AggregateBase) LoadEvents(aggregateId core.IIdentity) []IEvt {
	//TODO implement me
	panic("implement me")
}

func (a *AggregateBase) Attempt(cmd ICmd) (IFbk, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AggregateBase) Apply(evt IEvt) {
	//TODO implement me
	panic("implement me")
}

func (a *AggregateBase) LoadState(aggregateId core.IIdentity) interface{} {
	//TODO implement me
	panic("implement me")
}

func (a *AggregateBase) AppendEvent(evt IEvt) {
	//TODO implement me
	panic("implement me")
}

func NewAggregateBase(eventStream IEventStream, mediator mediator.IDECBus) IAggregate {
	return &AggregateBase{
		EventStream: eventStream,
		Mediator:    mediator}
}
