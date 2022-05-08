package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/domain"
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
)

type Aggregate struct {
	store sdk.IStore
	bus   sdk.IDECBus
	state *model.Greenhouse
}

func (a *Aggregate) Attempt(cmd sdk.ICmd) sdk.IFbk {
	fbk := NewFbk("", model.Error, "", "")
	if &cmd == nil {
		fbk.error = "cmd cannot be nil"
		return fbk
	}
	c := cmd.(*Cmd)
	id := c.AggregateId().(*sdk.Identity)
	a.state = infra.LoadGreenhouse(a.store, id.Id())
	if a.state == nil {
		a.state = model.NewGreenhouse(id.Value, c.Name, c.Settings, c.Sensor, c.Fan, c.Sprinkler)
		infra.SaveGreenhouse(a.store, a.state)
	}
	if !a.state.Status.HasFlag(model.Initialized) {
		evt := NewEvt(c.AggregateId(), c.TraceId(), c.Settings)
		a.Raise(evt)
		fbk = NewFbk(c.AggregateId().Id(), a.state.Status, c.TraceId(), "")
	}
	return fbk
}

func (a *Aggregate) Raise(evt *Evt) {
	a.Apply(evt)
	domain.PublishEvent(a.bus, EVT_TOPIC, evt)
}

func (a *Aggregate) Apply(evt sdk.IEvt) {
	e := evt.(*Evt)
	a.state = infra.LoadGreenhouse(a.store, e.AggregateId().Id())
	a.state.Status = model.Initialized
	infra.SaveGreenhouse(a.store, a.state)
}

func NewAggregate(store sdk.IStore, bus sdk.IDECBus) *Aggregate {
	return &Aggregate{
		bus:   bus,
		store: store,
	}
}
