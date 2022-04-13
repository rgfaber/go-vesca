package main

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/domain/measure"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"time"
)

type ISupervisor interface {
	Supervise()
}

type Supervisor struct {
	config *config.Config
	//sensorId *sdk.Identity
	state    *model.Root
	bus      dec.IDECBus
	features []domain.IFeature
}

func NewSupervisor(cfg *config.Config,
	bus dec.IDECBus,
	features []domain.IFeature) *Supervisor {
	state := model.NewRoot(cfg)
	return &Supervisor{
		config: cfg,
		//sensorId: sdk.NewIdentityFrom(config.GO_VESCA_TH_SENSOR_PREFIX, cfg.SensorId()),
		state:    state,
		features: features,
		bus:      bus,
	}
}

func (s *Supervisor) spawn(f domain.IFeature, m *model.Root) {
	go func(feature domain.IFeature, state *model.Root) {
		feature.Run()
	}(f, m)
}

func (s *Supervisor) Supervise() {
	for _, f := range s.features {
		s.spawn(f, s.state)
		time.Sleep(1 * time.Second)
	}
	//	s.Initialize()
	fmt.Printf("Sensor [%+v] is up!\n", s.state.ID.Id())
}

func (s *Supervisor) Initialize() {
	s.bus.Publish(initialize.CMD_TOPIC, initialize.NewCmd(15.0, 50.0))
}

func (s *Supervisor) measure() {
	for {
		fmt.Println("Triggering Measurement in 2s")
		time.Sleep(2 * time.Second)
		s.bus.Publish(measure.CMD_TOPIC, measure.NewCmd())
	}
}
