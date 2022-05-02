package main

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/domain/initialize"
	"github.com/rgfaber/go-vesca/greenhouse/domain/measure"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"time"
)

type ISupervisor interface {
	Supervise()
}

func ImplementsISuperviser(sup ISupervisor) bool {
	return true
}

type Supervisor struct {
	config   *config.Config
	bus      dec.IDECBus
	features []dec.IFeature
}

func NewSupervisor(cfg *config.Config,
	bus dec.IDECBus,
	features []dec.IFeature) *Supervisor {
	return &Supervisor{
		config: cfg,
		//sensorId: sdk.NewIdentityFrom(config.GO_VESCA_TH_SENSOR_PREFIX, Config.sensorId()),
		features: features,
		bus:      bus,
	}
}

func (s *Supervisor) spawn(f dec.IFeature) {
	go func(feature dec.IFeature) {
		feature.Run()
	}(f)
}

func (s *Supervisor) Supervise() {
	for _, f := range s.features {
		s.spawn(f)
		time.Sleep(1 * time.Second)
	}
	//	s.Initialize()
	fmt.Printf("Supervisor is up!")
}

func (s *Supervisor) Initialize() {
	id := model.NewGreenhouseID(s.config.SensorId())
	traceId, _ := sdk.NewUuid()
	cmd := initialize.NewCmd(*id, s.config.SensorName(), s.config.GreenhouseId(), traceId, 15.0, 42.0)
	s.bus.Publish(initialize.CMD_TOPIC, cmd)
}

func (s *Supervisor) measure() {
	for {
		fmt.Println("Triggering neasurement in 2s")
		time.Sleep(2 * time.Second)
		id := model.NewGreenhouseID(s.config.SensorId())
		traceId, _ := sdk.NewUuid()
		cmd := measure.NewCmd(*id, traceId)
		s.bus.Publish(measure.CMD_TOPIC, cmd)
	}
}
