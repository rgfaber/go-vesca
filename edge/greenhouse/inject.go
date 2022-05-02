package main

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/features/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
	initializeInfra "github.com/rgfaber/go-vesca/th-sensor/infra/initialize"
)

var (
	Config            = config.NewConfig()
	MemStore          = infra.NewStore()
	MemBus            = dec.NewDECBus()
	NatsBus           = infra.NewNatsBus(Config)
	InitializeEmitter = initializeInfra.NewEmitter(NatsBus)
	Features          = []dec.IFeature{
		initialize.NewFeature(MemBus, MemStore, InitializeEmitter),
		//kill.NewFeature(Root, MemBus),
		//measure.NewFeature(MemBus, MemStore),
	}
	TheSupervisor = NewSupervisor(Config, MemBus, Features)
)
