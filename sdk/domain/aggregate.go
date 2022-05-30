package domain

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
)

type IEventStream interface {
	LoadEvents(aggregateId core.IIdentity) error
}

type IAggregate interface {
	Attempt(cmd ICmd) IFbk
	Apply(evt IEvt)
	LoadEvents(aggregateId core.IIdentity) []IEvt
	AppendEvent(evt IEvt)
}

func ImplementsIAggregate(aggregate IAggregate) bool {
	return true
}

type AggregateBase struct {
	EventStream IEventStream
	Mediator    mediator.IDECBus
	State       interface{}
}

func (a *AggregateBase) LoadEvents(aggregateId core.IIdentity) []IEvt {
	//TODO implement me
	panic("implement me")
}

func (a *AggregateBase) Attempt(cmd ICmd) IFbk {
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
