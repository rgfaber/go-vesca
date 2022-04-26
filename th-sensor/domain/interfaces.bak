package domain

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type IStore interface {
	Load(id string) *model.Root
	Save(model model.Root)
}

type ICmd interface {
	AggregateId() sdk.IIdentity
}

type IEvt interface {
	AggregateId() sdk.IIdentity
}

type IFbk interface {
	AggregateId() sdk.IIdentity
	Status() model.SensorStatus
	IsSuccess() bool
	TraceId() string
}

type IAggregate interface {
	Attempt(cmd ICmd) (IFbk, error)
	Apply(evt IEvt)
}
