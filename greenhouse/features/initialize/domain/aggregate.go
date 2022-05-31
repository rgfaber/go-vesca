package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/infra/memory"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/topics"
	"github.com/rgfaber/go-vesca/greenhouse/helpers/domain"
	"github.com/rgfaber/go-vesca/greenhouse/helpers/infra"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core"
	sdk_domain "github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
)

var (
	Root = NewAggregate(memory.Store, memory.Mediator)
)

type Aggregate struct {
	store store.IStore
	bus   mediator.IDECBus
	state *model.Greenhouse
}

func (a *Aggregate) LoadEvents(aggregateId core.IIdentity) []sdk_domain.IEvt {
	//TODO implement me
	panic("implement me")
}

func (a *Aggregate) AppendEvent(evt sdk_domain.IEvt) {
	//TODO implement me
	panic("implement me")
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
		a.state = model.NewGreenhouse(id.Value, c.Details.Name, c.Settings, c.Sensor, c.Fan, c.Sprinkler)
		infra.SaveGreenhouse(a.store, a.state)
	}
	if !a.state.Status.HasFlag(model.Initialized) {
		evt := NewEvt(c.AggregateId(), c.TraceId(), a.state)
		a.Raise(evt)
		fbk = NewFbk(c.AggregateId().Id(), a.state.Status, c.TraceId(), "")
	}
	return fbk
}

func (a *Aggregate) Raise(evt *Evt) {
	a.Apply(evt)
	domain.PublishEvent(a.bus, topics.EVT_TOPIC, evt)
}

func (a *Aggregate) Apply(evt sdk_domain.IEvt) {
	e := evt.(*Evt)
	a.state = infra.LoadGreenhouse(a.store, e.AggregateId().Id())
	a.state.Status = model.Initialized
	infra.SaveGreenhouse(a.store, a.state)
}

func NewAggregate(store store.IStore, bus mediator.IDECBus) sdk_domain.IAggregate {
	return &Aggregate{
		bus:   bus,
		store: store,
	}
}
