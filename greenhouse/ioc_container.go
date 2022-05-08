package main

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize"
	nats2 "github.com/rgfaber/go-vesca/greenhouse/features/initialize/infra/nats"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/nats"
	sdk_nats_config "github.com/rgfaber/go-vesca/sdk/nats/config"
)

var (
	Config                = config.NewConfig()
	NatsConfig            = sdk_nats_config.NewNatsConfig()
	MemStore              = sdk.NewStore()
	MemBus                = sdk.NewDECBus()
	NatsBus               = nats.NewNatsBus(NatsConfig)
	InitializedSerializer = nats2.NewFactSerializer()
	InitializedEmitter    = nats2.NewEmitter(MemBus, NatsBus, nats2.FACT_TOPIC, InitializedSerializer)
	Features              = []sdk.IFeature{
		initialize.NewFeature(MemBus,
			MemStore,
			InitializedEmitter),
		//kill.NewFeature(Root, MemBus),
		//measure.NewFeature(MemBus, MemStore),
	}
)
