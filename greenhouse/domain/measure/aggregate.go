package measure

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/domain"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/dec"
)

type Aggregate struct {
	bus   dec.IDECBus
	store dec.IStore
	state *model.Greenhouse
}

func (a *Aggregate) Attempt(cmd dec.ICmd) (dec.IFbk, error) {
	if cmd == nil {
		return nil, fmt.Errorf("command cannot be nil")
	}
	c := cmd.(*Cmd)
	a.state = domain.LoadState(a.store, c.AggregateId().Id())
	if a.state == nil {
		return nil, fmt.Errorf("aggregate [%+v] not found", c.AggregateId().Id())
	}
	if !a.state.Status.HasFlag(model.Initialized) {
		return nil, fmt.Errorf("aggregate [%+v] is not initialized", c.AggregateId().Id())
	}
	temp := a.state.Settings.Temperature
	hum := a.state.Settings.Humidity
	if a.state.Status.HasFlag(model.FanActivated) {
		temp -= .2
	} else {
		temp += .3
	}
	if a.state.Status.HasFlag(model.SprinklerActivated) {
		hum += 2.0
	} else {
		hum -= 3.0
	}
	evt := NewEvt(c.AggregateId(), c.traceId, temp, hum)
	raise(a, evt)
	return NewFbk(c.AggregateId().Id(), c.traceId, true, a.state.Status), nil
}

func raise(a *Aggregate, e *Evt) {
	a.Apply(e)
	a.bus.Publish(EVT_TOPIC, e)
}

func (a *Aggregate) Apply(evt dec.IEvt) {
	e := evt.(*Evt)
	a.state = domain.LoadState(a.store, e.aggregateId.Id())
	a.state.Settings.Temperature = e.temperature
	a.state.Settings.Humidity = e.humidity
	domain.SaveState(a.store, a.state)

}

func NewAggregate(store dec.IStore, bus dec.IDECBus) *Aggregate {
	return &Aggregate{
		bus:   bus,
		store: store,
	}
}
