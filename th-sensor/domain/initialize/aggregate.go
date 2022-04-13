package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Aggregate struct {
	ID    sdk.IIdentity
	store domain.IStore
	bus   dec.IDECBus
	state model.Root
}

func (a *Aggregate) Execute(cmd domain.ICmd) (domain.IRsp, error) {
	if &cmd == nil {
		return nil, fmt.Errorf("initialize.Execute requires an initialize.Cmd")
	}
	a.state = a.store.Load()
	if !a.state.Status.HasFlag(model.Initialized) {
		evt := NewEvt(cmd.traceId, cmd.measurement)
		a.Raise(evt)
		return NewRsp(a.state.ID.Id(), cmd.traceId), nil
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
		ID:    identity,
		bus:   bus,
		store: store,
	}
}
