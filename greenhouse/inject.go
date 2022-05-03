package main

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize"
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	initializeInfra "github.com/rgfaber/go-vesca/greenhouse/infra/initialize"
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
