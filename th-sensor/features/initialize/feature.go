package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"log"
)

type Feature struct {
	bus       dec.IDECBus
	aggregate domain.IAggregate
}

func NewFeature(bus dec.IDECBus, store domain.IStore) *Feature {
	return &Feature{
		bus:       bus,
		aggregate: initialize.NewAggregate(store, bus),
	}
}

func (f *Feature) Raise(evt initialize.Evt) {
	f.aggregate.Apply(evt)
	f.bus.Publish(initialize.EVT_TOPIC, evt)
}

func (f *Feature) Listen() {
	err := f.bus.Subscribe(initialize.EVT_TOPIC, func(evt initialize.Evt) {
		f.apply(evt)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("initialize.Feature -> Listening for [%+v]\n", initialize.EVT_TOPIC)
}

func (f *Feature) Respond() {
	err := f.bus.Subscribe(initialize.CMD_TOPIC, func(cmd initialize.Cmd) {
		f.exec(cmd)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("initialize.Feature -> Responding to [%+v]\n", initialize.CMD_TOPIC)
}

func (f *Feature) Run() {
	f.Listen()
	f.Respond()
	fmt.Println("initialize.Feature is Up!")
}

func (f *Feature) apply(evt initialize.Evt) {
	f.state.Status.Set(model.Initialized)
}

func (f *Feature) exec(cmd *initialize.Cmd) {
	if cmd == nil {
		err := fmt.Errorf("Argument 'cmd' cannot be nil")
		panic(err)
	}
	if &cmd.measurement == nil {
		err := fmt.Errorf("Argument 'cmd' cannot be nil")
		panic(err)
	}
	if !f.state.Status.HasFlag(model.Initialized) {
		f.raise(initialize.NewEvt(cmd.measurement))
	}
}

func (f *Feature) raise(evt *initialize.Evt) {
	f.bus.Publish(initialize.EVT_TOPIC, evt)
}
