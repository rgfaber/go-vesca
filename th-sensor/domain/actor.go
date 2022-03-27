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
	Kill     *kill.Handler
}

func NewActor(sensorId string) *Actor {
	return &Actor{
		State:    model.NewRoot(sensorId),
		Handlers: make([]IHandler, 10),
		Kill:     kill.NewHandler(),
	}
}

func (*Actor) Register(handler IHandler) {
}

func (a *Actor) Exec(cmd kill.Cmd) {

}

func (*Actor) Apply(evt initialize.Evt) {
}

func (*Actor) Exec(cmd initialize.Cmd) {
}

func (a *Actor) Run() {

	for evt := range a.Kill.Events {
		go func(kill kill.Evt) {
			fmt.Println("Received Kill event, terminating in 5 seconds")
			time.Sleep(5 * time.Second)
			os.Exit(0)
		}(evt)
	}
}
