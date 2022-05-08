package main

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/domain"
	domain2 "github.com/rgfaber/go-vesca/greenhouse/features/measure/domain"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"time"
)

var TheSupervisor = NewSupervisor(Config, MemBus, Features)

type ISupervisor interface {
	Supervise()
}

func ImplementsISuperviser(sup ISupervisor) bool {
	return true
}

type Supervisor struct {
	config   *config.Config
	bus      sdk.IDECBus
	features []sdk.IFeature
}

func NewSupervisor(cfg *config.Config,
	bus sdk.IDECBus,
	features []sdk.IFeature) *Supervisor {
	return &Supervisor{
		config: cfg,
		//sensorId: sdk.NewIdentityFrom(config.GO_VESCA_TH_SENSOR_PREFIX, BogusConfig.sensorId()),
		features: features,
		bus:      bus,
	}
}

func (s *Supervisor) spawn(f sdk.IFeature) {
	go func(feature sdk.IFeature) {
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
	id := model.NewGreenhouseID(s.config.GreenhouseId())
	traceId, _ := sdk.NewUuid()
	settings := model.NewSettings(15.0, 42.0)
	sensor := model.NewSensor(s.config.SensorId(), s.config.SensorName())
	fan := model.NewFan(s.config.FanId(), s.config.FanName())
	sprinkler := model.NewSprinkler(s.config.SprinklerId(), s.config.SprinklerName())
	cmd := domain.NewCmd(id, traceId, s.config.GreenhouseName(), settings, sensor, fan, sprinkler)
	s.bus.Publish(domain.CMD_TOPIC, cmd)
}

func (s *Supervisor) measure() {
	for {
		fmt.Println("Triggering neasurement in 2s")
		time.Sleep(2 * time.Second)
		id := model.NewGreenhouseID(s.config.SensorId())
		traceId, _ := sdk.NewUuid()
		cmd := domain2.NewCmd(*id, traceId)
		s.bus.Publish(domain2.CMD_TOPIC, cmd)
	}
}
