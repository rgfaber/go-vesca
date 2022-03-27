package domain

import (
	"fmt"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/domain/kill"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"os"
	"time"
)

type IHandler interface {
	Listen()
}

type Actor struct {
	State    *model.Root
	Handlers []IHandler
	Killer   *kill.Handler
}

func NewActor(sensorId string) *Actor {
	return &Actor{
		State:    model.NewRoot(sensorId),
		Handlers: make([]IHandler, 10),
		Killer:   kill.NewHandler(),
	}
}

func (*Actor) Register(handler IHandler) {
}

func (a *Actor) Kill() {
	cmd := kill.NewCnd()
	a.Killer.Commands <- *cmd
}

func (*Actor) Apply(evt initialize.Evt) {
}

func (*Actor) Exec(cmd initialize.Cmd) {
}

func (a *Actor) Run() {
	fmt.Println("Sensor with Id", a.State.Id.Id(), "is Up!")
	a.Killer.Run()

	for evt := range a.Killer.Events {
		go func(kill kill.Evt) {
			fmt.Println("Received Kill event, terminating in 5 seconds")
			time.Sleep(5 * time.Second)
			os.Exit(0)
		}(evt)
	}
}
