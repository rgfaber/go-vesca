package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"log"
)

type Feature struct {
	bus       dec.IDECBus
	aggregate dec.IAggregate
}

func (f *Feature) Bus() dec.IDECBus {
	return f.bus
}

func (f *Feature) Store() dec.IStore {
	return f.Store()
}

func (f *Feature) Aggregate() dec.IAggregate {
	return f.aggregate
}

func NewFeature(bus dec.IDECBus, store dec.IStore) *Feature {
	return &Feature{
		bus:       bus,
		aggregate: initialize.NewAggregate(store, bus),
	}
}

func (f *Feature) StartListening() {
	fmt.Printf("initialize.Feature -> not listening for any domain events")
}

func (f *Feature) StartResponding() {
	err := f.bus.Subscribe(initialize.CMD_TOPIC, func(cmd *initialize.Cmd) {
		f.aggregate.Attempt(cmd)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("initialize.Feature -> Responding to [%+v]\n", initialize.CMD_TOPIC)
}

func (f *Feature) Run() {
	f.StartListening()
	f.StartResponding()
	fmt.Println("initialize.Feature is Up!")
}
