package domain

import "github.com/rgfaber/go-vesca/th-sensor/model"

type IStore interface {
	Load() model.Root
	Save(model model.Root)
}

type ICmd interface{}

type IEvt interface{}

type IRsp interface{}

type IAggregate interface {
	Execute(cmd ICmd) (IRsp, error)
	Apply(evt IEvt)
}
