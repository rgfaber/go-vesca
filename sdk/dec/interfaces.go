package dec

import (
	"github.com/rgfaber/go-vesca/sdk"
)

type IStore interface {
	Load(id string) interface{}
	Save(model interface{})
}

type ICmd interface {
	AggregateId() sdk.IIdentity
}

type IEvt interface {
	AggregateId() sdk.IIdentity
}

type IFbk interface {
	AggregateId() sdk.IIdentity
	Status() int
	IsSuccess() bool
	TraceId() string
}

type IAggregate interface {
	Attempt(cmd ICmd) (IFbk, error)
	Apply(evt IEvt)
}
