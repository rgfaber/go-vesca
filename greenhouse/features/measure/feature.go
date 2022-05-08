package measure

import (
	"fmt"
	domain2 "github.com/rgfaber/go-vesca/greenhouse/features/measure/domain"
	"github.com/rgfaber/go-vesca/sdk"
	"log"
)

type Feature struct {
	bus       sdk.IDECBus
	store     sdk.IStore
	aggregate sdk.IAggregate
}

func (f *Feature) Listen() {
	fmt.Printf("measure.Feature -> Does not listen for any domain events")
}

func (f *Feature) Respond() {
	err := f.bus.Subscribe(domain2.CMD_TOPIC, func(cmd sdk.ICmd) {
		fmt.Printf("measure.BogusCmd  => responds to [%+v]", domain2.CMD_TOPIC)
		f.aggregate = domain2.NewAggregate(f.store, f.bus)
		f.aggregate.Attempt(cmd)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("measure.Feature -> Responding to [%+v]\n", domain2.CMD_TOPIC)
}

func NewFeature(bus sdk.IDECBus, store sdk.IStore) *Feature {
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
