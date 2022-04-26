package measure

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/domain/measure"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"log"
)

type Feature struct {
	state *model.Root
	bus   dec.IDECBus
}

func (f *Feature) Listen() {
	err := f.bus.Subscribe(measure.EVT_TOPIC, func(evt *measure.Evt) {
		f.apply(evt)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("measure.Feature -> Listening for [%+v]\n", measure.EVT_TOPIC)
}

func (f *Feature) Respond() {
	err := f.bus.Subscribe(measure.CMD_TOPIC, func(cmd measure.Cmd) {
		fmt.Printf("Received measure.Cmd on [%+v]", measure.CMD_TOPIC)
		f.exec()
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("measure.Feature -> Responding to [%+v]\n", measure.CMD_TOPIC)
}

func NewFeature(state *model.Root, bus dec.IDECBus) *Feature {
	return &Feature{
		state: state,
		bus:   bus,
	}
}

func (f *Feature) Run() {
	go f.Listen()
	go f.Respond()
	fmt.Println("measure.Feature is Up!")
}
