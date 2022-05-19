package domain

import (
	"github.com/rgfaber/go-vesca/go-scream/core"
	mem_interfaces "github.com/rgfaber/go-vesca/go-scream/infra/memory/interfaces"
	"github.com/rgfaber/go-vesca/go-scream/interfaces"
	"github.com/rgfaber/go-vesca/greenhouse/helpers/domain"
	"github.com/rgfaber/go-vesca/greenhouse/helpers/infra"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/infra/memory"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/topics"
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

type Aggregate struct {
	store mem_interfaces.IStore
	bus   mem_interfaces.IDECBus
	state *model.Greenhouse
}

func (a *Aggregate) Attempt(cmd interfaces.ICmd) interfaces.IFbk {
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

func (a *Aggregate) Apply(evt interfaces.IEvt) {
	e := evt.(*Evt)
	a.state = infra.LoadGreenhouse(a.store, e.AggregateId().Id())
	a.state.Status = model.Initialized
	infra.SaveGreenhouse(a.store, a.state)
}

func NewAggregate(store mem_interfaces.IStore, bus mem_interfaces.IDECBus) *Aggregate {
	return &Aggregate{
		bus:   bus,
		store: store,
	}
}
