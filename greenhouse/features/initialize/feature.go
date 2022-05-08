package initialize

import (
	"fmt"
	domain2 "github.com/rgfaber/go-vesca/greenhouse/features/initialize/domain"
	"github.com/rgfaber/go-vesca/sdk"
	"log"
)

type Feature struct {
	memBus    sdk.IDECBus
	aggregate sdk.IAggregate
	emitter   sdk.IFactEmitter
}

func (f *Feature) Bus() sdk.IDECBus {
	return f.memBus
}

func (f *Feature) Store() sdk.IStore {
	return f.Store()
}

func (f *Feature) Aggregate() sdk.IAggregate {
	return f.aggregate
}

func NewFeature(bus sdk.IDECBus,
	store sdk.IStore,
	emitter sdk.IFactEmitter) *Feature {
	return &Feature{
		memBus:    bus,
		aggregate: domain2.NewAggregate(store, bus),
		emitter:   emitter,
	}
}

func (f *Feature) HandleDomainEvents() {
	err := f.memBus.Subscribe(domain2.EVT_TOPIC, func(evt *domain2.Evt) {
		//		fact := initialize.NewFact(evt.aggregateId().aggregateId(), evt.)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("initialize.Feature -> listening on [%+v]", domain2.EVT_TOPIC)
}

func (f *Feature) HandleCommand() {
	err := f.memBus.Subscribe(domain2.CMD_TOPIC, func(cmd *domain2.Cmd) {
		f.aggregate.Attempt(cmd)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("initialize.Feature -> Responding to [%+v]\n", domain2.CMD_TOPIC)
}

func (f *Feature) Run() {
	f.HandleDomainEvents()
	f.HandleCommand()
	fmt.Println("initialize.Feature is Up!")
}
