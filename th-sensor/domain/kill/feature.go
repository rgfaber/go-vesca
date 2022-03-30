package kill

import (
	"fmt"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"log"
	"os"
)

type Feature struct {
	state *model.Root
	chCmd chan *Cmd
	chEvt chan *Evt
}

func NewFeature(state *model.Root, chCmd ChCmd, chEvt ChEvt) *Feature {
	return &Feature{
		state: state,
		chCmd: chCmd,
		chEvt: chEvt,
	}
}

func (a *Feature) Exec(cmd *Cmd) {
	if cmd == nil {
		err := fmt.Errorf("Command cannot be nil!")
		panic(err)
	}
	if a.state.Status.HasFlag(model.Created) {
		a.Raise(cmd.ToEvt())
	}
}

func (a *Feature) Raise(evt *Evt) {
	a.chEvt <- evt
}

func (a *Feature) Apply(evt *Evt) {
	if evt == nil {
		log.Fatalf("Received NIL kill.Evt")
	}
	a.state.Status.Set(model.Killed)
	os.Exit(0)
}

func (a *Feature) Listen() {
	for evt := range a.chEvt {
		go func(e *Evt) {
			a.Apply(e)
		}(evt)
	}
}

func (a *Feature) Respond() {
	for cmd := range a.chCmd {
		go func(c *Cmd) {
			a.Exec(c)
		}(cmd)
	}
}

func (a *Feature) Run() {
	fmt.Println("kill.Feature is Up!")
	go a.Listen()
	go a.Respond()
	select {}
}
