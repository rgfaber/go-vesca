package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"log"
)

type Feature struct {
	state *model.Root
	bus   dec.IDECBus
}

func NewFeature(r *model.Root, bus dec.IDECBus) *Feature {
	return &Feature{
		bus:   bus,
		state: r,
	}
}

func (f *Feature) Listen() {
	err := f.bus.Subscribe(EVT_TOPIC, func(evt *Evt) {
		f.apply(evt)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("initialize.Feature -> Listening for [%+v]\n", EVT_TOPIC)
}

func (f *Feature) Respond() {
	err := f.bus.Subscribe(CMD_TOPIC, func(cmd *Cmd) {
		f.exec(cmd)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("initialize.Feature -> Responding to [%+v]\n", CMD_TOPIC)
}

func (f *Feature) Run() {
	f.Listen()
	f.Respond()
	fmt.Println("initialize.Feature is Up!")
}

func (f *Feature) apply(evt *Evt) {
	f.state.Status.Set(model.Initialized)
}

func (f *Feature) exec(cmd *Cmd) {
	if cmd == nil {
		return
	}
	if !f.state.Status.HasFlag(model.Initialized) {
		f.raise(NewEvt())
	}
}

func (f *Feature) raise(evt *Evt) {
	f.bus.Publish(EVT_TOPIC, evt)
}
