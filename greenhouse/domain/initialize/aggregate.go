package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/domain"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
)

type Aggregate struct {
	store dec.IStore
	bus   dec.IDECBus
	state *model.Greenhouse
}

func (a *Aggregate) Attempt(cmd dec.ICmd) (dec.IFbk, error) {
	if &cmd == nil {
		return nil, fmt.Errorf("initialize.Attempt requires an initialize.Cmd")
	}
	c := cmd.(*Cmd)
	id := c.AggregateId().(*sdk.Identity)
	a.state = domain.LoadState(a.store, id.Id())
	if a.state == nil {
		a.state = model.NewGreenhouse(id.Value(), c.name, c.settings, c.sensor, c.fan, c.sprinkler)
		a.store.Save(*a.state)
	}
	if !a.state.Status.HasFlag(model.Initialized) {
		evt := NewEvt(c.AggregateId(), c.traceId, c.settings)
		a.Raise(evt)
		return NewFbk(c.AggregateId().Id(), a.state.Status, true, c.traceId), nil
	}
	return nil, fmt.Errorf("Aggregate [%+v] has already been initialized", a.state.ID.Id())
}

func (a *Aggregate) Raise(evt *Evt) {
	a.Apply(evt)
	domain.Publish(a.bus, EVT_TOPIC, evt)
}

func (a *Aggregate) Apply(evt dec.IEvt) {
	e := evt.(*Evt)
	a.state = domain.LoadState(a.store, e.AggregateId().Id())
	a.state.Status = model.Initialized
	domain.SaveState(a.store, a.state)
	a.store.Save(*a.state)
}

func NewAggregate(store dec.IStore, bus dec.IDECBus) *Aggregate {
	return &Aggregate{
		bus:   bus,
		store: store,
	}
}
