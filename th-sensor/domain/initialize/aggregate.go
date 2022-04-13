package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Aggregate struct {
	store domain.IStore
	bus   dec.IDECBus
	state model.Root
}

func (a *Aggregate) Execute(cmd domain.ICmd) (domain.IRsp, error) {
	if &cmd == nil {
		return nil, fmt.Errorf("initialize.Execute requires an initialize.Cmd")
	}
	c := cmd.(Cmd)
	a.state = a.store.Load(c.aggregateId)
	if !a.state.Status.HasFlag(model.Initialized) {
		evt := NewEvt(c.aggregateId, c.traceId, c.measurement)
		a.Raise(evt)
		return NewRsp(a.state.ID, cmd.traceId), nil
	}
	return nil, fmt.Errorf("Aggregate [%+v] has already been initialized", a.state.ID.Id())
}

func (a *Aggregate) Raise(evt *Evt) {
	a.bus.Publish(EVT_TOPIC, *evt)
}

func (a *Aggregate) Apply(evt domain.IEvt) {
	e := evt.(Evt)
	m := a.store.Load(e.aggregateId)
	a.state.Status = model.Initialized
	a.store.Save(a.state)
}

func NewAggregate(store domain.IStore, bus dec.IDECBus) *Aggregate {
	return &Aggregate{
		bus:   bus,
		store: store,
	}
}
