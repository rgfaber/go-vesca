package kill

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"log"
	"os"
)

type Feature struct {
	state *model.Root
	bus   dec.IDECBus
}

func NewFeature(state *model.Root, bus dec.IDECBus) *Feature {
	return &Feature{
		state: state,
		bus:   bus,
	}
}

func (a *Feature) exec(cmd *Cmd) {
	if cmd == nil {
		err := fmt.Errorf("Command cannot be nil!")
		panic(err)
	}
	if a.state.Status.HasFlag(model.Created) {
		a.raise(cmd.ToEvt())
	}
}

func (a *Feature) raise(evt *Evt) {
	fmt.Printf("kill.Feature -> Raising [%+v] on [%+v]\n", evt, EVT_TOPIC)
	a.bus.Publish(EVT_TOPIC, NewEvt())
}

func (a *Feature) apply(evt *Evt) {
	if evt == nil {
		log.Fatalf("Received NIL kill.Evt")
	}
	a.state.Status.Set(model.Killed)
	os.Exit(0)
}

func (a *Feature) Listen() {
	err := a.bus.Subscribe(EVT_TOPIC, func(evt *Evt) {
		a.apply(evt)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("kill.Feature -> Listening for [%+v]\n", EVT_TOPIC)
}

func (a *Feature) Respond() {
	err := a.bus.Subscribe(CMD_TOPIC, func(cmd *Cmd) {
		a.exec(cmd)
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("kill.Feature -> Responding to [%+v]\n", CMD_TOPIC)
}

func (a *Feature) Run() {
	a.Listen()
	a.Respond()
	fmt.Println("kill.Feature is Up!")
	select {}
}
