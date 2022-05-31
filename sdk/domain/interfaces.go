package domain

import "github.com/rgfaber/go-vesca/sdk/core"

type IEventStream interface {
	LoadEvents(aggregateId core.IIdentity) error
	AppendEvent(aggregateId core.IIdentity) error
	Broadcast(events []IEvt) error
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
