package main

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize"
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	initializeInfra "github.com/rgfaber/go-vesca/greenhouse/infra/initialize/nats"
)

var (
	Config            = config.NewConfig()
	MemStore          = infra.NewStore()
	MemBus            = sdk.NewDECBus()
	NatsBus           = infra.NewNatsBus(Config)
	InitializeEmitter = initializeInfra.NewEmitter(NatsBus)
	Features          = []sdk.IFeature{
		initialize.NewFeature(MemBus, MemStore, InitializeEmitter),
		//kill.NewFeature(Root, MemBus),
		//measure.NewFeature(MemBus, MemStore),
	}
	TheSupervisor = NewSupervisor(Config, MemBus, Features)
)
