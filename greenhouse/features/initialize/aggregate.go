package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/helpers/domain"
	"github.com/rgfaber/go-vesca/greenhouse/helpers/infra"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core"
	mem_interfaces "github.com/rgfaber/go-vesca/sdk/infra/memory/interfaces"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type Aggregate struct {
	store mem_interfaces.IStore
	bus   mem_interfaces.IDECBus
	state *model.Greenhouse
}

func (a *Aggregate) Attempt(cmd interfaces.ICmd) interfaces.IFbk {
	fbk := NewFbk("", model.Error, "", "")
	if &cmd == nil {
		fbk.error = "cmd cannot be nil"
		return fbk
	}
	c := cmd.(*Cmd)
	id := c.AggregateId().(*core.Identity)
	a.state = infra.LoadGreenhouse(a.store, id.Id())
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
	domain.PublishEvent(a.bus, EVT_TOPIC, evt)
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
