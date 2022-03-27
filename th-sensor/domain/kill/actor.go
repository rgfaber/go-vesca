package kill

import (
	"fmt"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Actor struct {
	State    *model.Root
	Commands chan *Cmd
	Events   chan *Evt
}

func NewActor(state *model.Root) *Actor {
	return &Actor{
		State:    state,
		Commands: make(chan *Cmd, 100),
		Events:   make(chan *Evt, 100)}
}

func (a *Actor) Exec(cmd *Cmd) {
	if cmd == nil {
		err := fmt.Errorf("Command cannot be nil!")
		panic(err)
	}
	a.Raise(cmd.ToEvt())
}

func (a *Actor) Raise(evt *Evt) {
	a.Events <- evt
}

func (a *Actor) Handle(evt *Evt) {

}

func (a *Actor) Listen() {
	for evt := range a.Events {
		go func(e *Evt) {
			a.Handle(e)
		}(evt)
	}
}

func (a *Actor) Respond() {
	for cmd := range a.Commands {
		go func(c Cmd) {
			a.Exec(c)
		}(cmd)
	}
}

func (a *Actor) Run() {
	fmt.Println("kill.Actor is Up!")
	go a.Listen()
	go a.Respond()
}

