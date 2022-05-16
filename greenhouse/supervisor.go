package main

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize"
	"github.com/rgfaber/go-vesca/greenhouse/features/measure"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core"
	memory_mediator "github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	sdk_interfaces "github.com/rgfaber/go-vesca/sdk/interfaces"
	"sync"
	"time"
)

var (
	cfg           = config.NewConfig()
	mediator      = memory_mediator.SingletonDECBus()
	TheSupervisor = NewSupervisor(initialize.Feature)
)

var lock = &sync.Mutex{}

var singleSup *Supervisor

type ISupervisor interface {
	Supervise()
}

func ImplementsISupervisor(sup ISupervisor) bool {
	return true
}

type Supervisor struct {
	features []sdk_interfaces.IFeature
}

func NewSupervisor(features ...sdk_interfaces.IFeature) *Supervisor {
	if singleSup == nil {
		lock.Lock()
		defer lock.Unlock()
		singleSup = &Supervisor{
			features: features,
		}
	}
	return singleSup
}

func (s *Supervisor) spawn(f sdk_interfaces.IFeature) {
	go func(feature sdk_interfaces.IFeature) {
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
	id := model.NewGreenhouseID(cfg.GreenhouseId())
	traceId, _ := core.NewUuid()
	settings := model.NewSettings(15.0, 42.0)
	sensor := model.NewSensor(cfg.SensorId(), cfg.SensorName())
	fan := model.NewFan(cfg.FanId(), cfg.FanName())
	sprinkler := model.NewSprinkler(cfg.SprinklerId(), cfg.SprinklerName())
	details := model.NewDetails(cfg.GreenhouseName(), "")
	cmd := initialize.NewCmd(id, traceId, details, settings, sensor, fan, sprinkler)
	mediator.Publish(initialize.CMD_TOPIC, cmd)
}

func (s *Supervisor) measure() {
	for {
		fmt.Println("Triggering neasurement in 2s")
		time.Sleep(2 * time.Second)
		id := model.NewGreenhouseID(cfg.SensorId())
		traceId, _ := core.NewUuid()
		cmd := measure.NewCmd(*id, traceId)
		mediator.Publish(measure.CMD_TOPIC, cmd)
	}
}
