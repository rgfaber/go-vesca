package measure

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/domain/measure"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"log"
)

type Feature struct {
	bus       dec.IDECBus
	store     dec.IStore
	aggregate dec.IAggregate
}

func (f *Feature) Listen() {
	fmt.Printf("measure.Feature -> Does not listen for any domain events")
}

func (f *Feature) Respond() {
	err := f.bus.Subscribe(measure.CMD_TOPIC, func(cmd dec.ICmd) {
		fmt.Printf("measure.Cmd  => responds to [%+v]", measure.CMD_TOPIC)
		f.aggregate = measure.NewAggregate(f.store, f.bus)
		f.aggregate.Attempt(cmd)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("measure.Feature -> Responding to [%+v]\n", measure.CMD_TOPIC)
}

func NewFeature(bus dec.IDECBus, store dec.IStore) *Feature {
	return &Feature{
		store: store,
		bus:   bus,
	}
}

func (f *Feature) Run() {
	go f.Listen()
	go f.Respond()
	fmt.Println("measure.Feature is Up!")
}
