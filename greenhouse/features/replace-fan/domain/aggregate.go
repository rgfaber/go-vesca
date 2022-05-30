package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/features/replace-fan/infra/memory"
	"github.com/rgfaber/go-vesca/greenhouse/features/replace-fan/topics"
	"github.com/rgfaber/go-vesca/greenhouse/helpers/domain"
	"github.com/rgfaber/go-vesca/greenhouse/helpers/infra"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core"
	sdk_domain "github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
)

type Aggregate struct {
	store store.IStore
	bus   mediator.IDECBus
	state *model.Greenhouse
}

func (a *Aggregate) Attempt(cmd sdk_domain.ICmd) sdk_domain.IFbk {
	fbk := NewFbk("", model.Error, "", "")
	if &cmd == nil {
		fbk.Err = "cmd cannot be nil"
		return fbk
	}
	c := cmd.(*Cmd)
	id := c.AggregateId().(*core.Identity)
	a.state = infra.LoadGreenhouse(memory.Store, id.Id())
	if a.state == nil {
		fbk.Err = "Greenhouse [%+v] not found."
		return fbk
	}
	if !a.state.Status.HasFlag(model.Initialized) {
		fbk.Err = "Greenhouse [%+v] not Initialized."
		return fbk
	}
	evt := NewEvt(id, c.TraceId(), c.Fan)
	a.Raise(evt)
	return fbk
}

func (a *Aggregate) Raise(evt *Evt) {
	a.Apply(evt)
	domain.PublishEvent(a.bus, topics.EVT_TOPIC, evt)
}

func (a *Aggregate) Apply(evt sdk_domain.IEvt) {
	e := evt.(*Evt)
	a.state = infra.LoadGreenhouse(a.store, e.AggregateId().Id())
	a.state.Status = model.FanReplaced
	infra.SaveGreenhouse(a.store, a.state)
}

func NewAggregate(store store.IStore, bus mediator.IDECBus) *Aggregate {
	return &Aggregate{
		bus:   bus,
		store: store,
	}
}
