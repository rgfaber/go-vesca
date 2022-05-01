package main

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/domain/measure"
	"github.com/rgfaber/go-vesca/th-sensor/features"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"time"
)

type ISupervisor interface {
	Supervise()
}

func ImplementsISuperviser(sup ISupervisor) bool {
	return true
}

type Supervisor struct {
	config *config.Config
	bus      dec.IDECBus
	features []features.IFeature
}

func NewSupervisor(cfg *config.Config,
	bus dec.IDECBus,
	features []features.IFeature) *Supervisor {
	return &Supervisor{
		config: cfg,
		//sensorId: sdk.NewIdentityFrom(config.GO_VESCA_TH_SENSOR_PREFIX, cfg.sensorId()),
		features: features,
		bus:      bus,
	}
}

func (s *Supervisor) spawn(f features.IFeature) {
	go func(feature features.IFeature) {
		feature.Run()
	}(f)
}

func (s *Supervisor) Supervise() {
	for _, f := range s.features {
		s.spawn(f)
		time.Sleep(1 * time.Second)
	}
	//	s.Initialize()
	fmt.Printf("Sensor [%+v] is up!\n", s.Id())
}

func (s *Supervisor) Initialize() {
	id := model.NewTHSensorId(s.config.SensorId())
	cmd := initialize.NewCmd(id, s.config.SensorName(), s.config.GreenhouseId(), sdk. )
	s.bus.Publish(initialize.CMD_TOPIC, initialize.NewCmd(15.0, 50.0))
}

func (s *Supervisor) measure() {
	for {
		fmt.Println("Triggering neasurement in 2s")
		time.Sleep(2 * time.Second)
		s.bus.Publish(measure.CMD_TOPIC, measure.NewCmd())
	}
}
