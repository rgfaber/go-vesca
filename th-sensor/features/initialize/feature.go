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

func NewFeature(bus dec.IDECBus, store dec.IStore) *Feature {
	return &Feature{
		bus:       bus,
		aggregate: initialize.NewAggregate(store, bus),
	}
}

func (f *Feature) Listen() {
	err := f.bus.Subscribe(initialize.EVT_TOPIC, func(evt initialize.Evt) {

	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("initialize.Feature -> Listening for [%+v]\n", initialize.EVT_TOPIC)
}

func (f *Feature) Respond() {
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
	f.Listen()
	f.Respond()
	fmt.Println("initialize.Feature is Up!")
}
