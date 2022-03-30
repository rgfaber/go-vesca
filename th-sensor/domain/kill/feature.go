package kill

import (
	"fmt"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Feature struct {
	State    *model.Root
	Commands chan *Cmd
	Events   chan *Evt
}

func NewFeature(state *model.Root) *Feature {
	return &Feature{
		State:    state,
		Commands: make(chan *Cmd, 100),
		Events:   make(chan *Evt, 100)}
}

func (a *Feature) Exec(cmd *Cmd) {
	if cmd == nil {
		err := fmt.Errorf("Command cannot be nil!")
		panic(err)
	}
	a.Raise(cmd.ToEvt())
}

func (a *Feature) Raise(evt *Evt) {
	a.Events <- evt
}

func (a *Feature) Handle(evt *Evt) {}

func (a *Feature) Listen() {
	for evt := range a.Events {
		go func(e *Evt) {
			a.Handle(e)
		}(evt)
	}
}

func (a *Feature) Respond() {
	for cmd := range a.Commands {
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