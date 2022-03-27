package domain

import (
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Actor struct {
	State      *model.Root
	Initialize *initialize.Handler
}

func NewActor(sensorId string) *Actor {
	return &Actor{
		State: model.NewRoot(sensorId),
	}
}

func (*Actor) Exec(cmd initialize.Cmd) {
}

func (*Actor) Exec(cnd activate.Cmd) {}

func (*Actor) Run() {

}
