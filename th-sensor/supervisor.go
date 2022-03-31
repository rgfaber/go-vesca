package main

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type ISupervisor interface {
}

type Supervisor struct {
	sensorId *sdk.Identity
	state    *model.Root
	mediator Mediator
	features []domain.IFeature
}

func NewSupervisor(cfg *config.Config,
	mediator Mediator,
	features []domain.IFeature) *Supervisor {
	state := model.NewRoot(cfg)
	return &Supervisor{
		sensorId: sdk.NewIdentityFrom(config.GO_VESCA_TH_SENSOR_PREFIX, cfg.SensorId()),
		state:    state,
		features: features,
		mediator: mediator,
	}
}

func spawn(f domain.IFeature, s *model.Root) {
	go func(feature domain.IFeature, state *model.Root) {
		feature.Run()
	}(f, s)
}

func (s *Supervisor) Supervise() {
	for _, f := range s.features {
		spawn(f, s.state)
	}
	fmt.Printf("Sensor [%+v] is up!\n", s.state.ID.Id())
}
