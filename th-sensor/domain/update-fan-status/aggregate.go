package update_fan_status

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Aggregate struct {
	bus   dec.IDECBus
	store dec.IStore
	state *model.Root
}

func (a *Aggregate) Attempt(cmd dec.ICmd) (dec.IFbk, error) {
	if cmd == nil {
		return nil, fmt.Errorf("Command must not be nil")
	}
	c := cmd.(*Cmd)
	id := cmd.AggregateId().Id()
	a.state = domain.LoadState(a.store, id)
	if a.state == nil {
		return nil, fmt.Errorf("[%+v] not found!", id)
	}
	if a.state.Status.HasFlag(model.Initialized) && (a.state.FanStatus != c.newStatus) {
		evt := NewEvt(c.sensorId, c.newStatus)
		a.raise(evt)
	}
	return NewFbk(c.sensorId, "", true, a.state.FanStatus), nil
}

func (a *Aggregate) raise(evt *Evt) {
	a.Apply(evt)
	domain.Publish(a.bus, EVT_TOPIC, evt)
}

func (a *Aggregate) Apply(evt dec.IEvt) {
	if evt == nil {
		panic("Event cannot be nil!")
	}
	e := evt.(*Evt)
	id := evt.AggregateId().Id()
	a.state = domain.LoadState(a.store, id)
	if a.state == nil {
		panic(fmt.Errorf("could not find [%+v] in store", id))
	}
	a.state.FanStatus = e.newStatus
	domain.SaveState(a.store, a.state)
}

func NewAggregate(store dec.IStore, bus dec.IDECBus) *Aggregate {
	return &Aggregate{
		bus:   bus,
		store: store,
	}
}
