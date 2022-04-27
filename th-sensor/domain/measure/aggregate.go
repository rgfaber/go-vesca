package measure

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Aggregate struct {
	bus   dec.IDECBus
	store dec.IStore
	state *model.Root
}

func (a *Aggregate) Attempt(cmd dec.ICmd) (dec.IFbk, error) {

	c := cmd.(*Cmd)

}

func (a *Aggregate) Apply(evt dec.IEvt) {
	//TODO implement me
	panic("implement me")
}

func NewAggregate(store dec.IStore, bus dec.IDECBus) *Aggregate {
	return &Aggregate{
		bus:   bus,
		store: store,
	}
}
