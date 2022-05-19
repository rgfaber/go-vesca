package measure

import (
	"fmt"
	"github.com/rgfaber/go-vesca/go-scream/infra/memory/interfaces"
	interfaces2 "github.com/rgfaber/go-vesca/go-scream/interfaces"
	"github.com/rgfaber/go-vesca/greenhouse/helpers/infra"
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

type Aggregate struct {
	bus   interfaces.IDECBus
	store interfaces.IStore
	state *model.Greenhouse
}

func (a *Aggregate) Attempt(cmd interfaces2.ICmd) (interfaces2.IFbk, error) {
	if cmd == nil {
		return nil, fmt.Errorf("command cannot be nil")
	}
	c := cmd.(*Cmd)
	a.state = infra.LoadGreenhouse(a.store, c.AggregateId().Id())
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
	return NewFbk(c.AggregateId().Id(), c.traceId, a.state.Status, ""), nil
}

func raise(a *Aggregate, e *Evt) {
	a.Apply(e)
	a.bus.Publish(EVT_TOPIC, e)
}

func (a *Aggregate) Apply(evt interfaces2.IEvt) {
	e := evt.(*Evt)
	a.state = infra.LoadGreenhouse(a.store, e.aggregateId.Id())
	a.state.Settings.Temperature = e.temperature
	a.state.Settings.Humidity = e.humidity
	infra.SaveGreenhouse(a.store, a.state)

}

func NewAggregate(store interfaces.IStore, bus interfaces.IDECBus) *Aggregate {
	return &Aggregate{
		bus:   bus,
		store: store,
	}
}
