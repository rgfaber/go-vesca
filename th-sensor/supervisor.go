package main

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/domain/kill"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"os"
	"time"
)

type ISupervisor interface {
}

type Supervisor struct {
	SensorId *sdk.Identity
	State    *model.Root
	Killer   *kill.Feature
	Features []domain.IFeature
}

func NewSupervisor(cfg *config.Config, features []domain.IFeature) *Supervisor {
	state := injectState(cfg)
	return &Supervisor{
		SensorId: sdk.NewIdentityFrom(config.GO_VESCA_TH_SENSOR_PREFIX, cfg.SensorId()),
		State:    state,
		Killer:   kill.NewFeature(state),
		Features: features,
	}
}

func injectState(cfg *config.Config) *model.Root {
	return model.NewRoot(cfg)
}

func (s *Supervisor) Kill() {
	cmd := kill.NewCnd("")
	s.Killer.Commands <- cmd
}

func (s *Supervisor) Supervise() {
	for _, f := range s.Features {
		go func(feature domain.IFeature, state *model.Root) {
			feature.Run()
		}(f, s.State)
	}
	fmt.Printf("Sensor [%+v] is up!\n", s.State.ID.Id())
	s.Killer.Run()
	for evt := range s.Killer.Events {
		go func(kill *kill.Evt) {
			fmt.Println("Received Kill event, terminating in 5 seconds")
			time.Sleep(5 * time.Second)
			os.Exit(0)
		}(evt)
	}
}
