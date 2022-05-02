package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/domain/initialize"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"log"
)

type Feature struct {
	bus       dec.IDECBus
	aggregate dec.IAggregate
	emitter   dec.IEmitter
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

func NewFeature(bus dec.IDECBus,
	store dec.IStore,
	emitter dec.IEmitter) *Feature {
	return &Feature{
		bus:       bus,
		aggregate: initialize.NewAggregate(store, bus),
		emitter:   emitter,
	}
}

func (f *Feature) StartListening() {
	err := f.bus.Subscribe(initialize.EVT_TOPIC, func(evt *initialize.Evt) {
		//		fact := initialize.NewFact(evt.AggregateId().Id(), evt.)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("initialize.Feature -> listening on [%+v]", initialize.EVT_TOPIC)
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
