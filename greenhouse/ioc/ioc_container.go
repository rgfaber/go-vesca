package ioc

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	initialize2 "github.com/rgfaber/go-vesca/greenhouse/features/initialize"
	nats2 "github.com/rgfaber/go-vesca/greenhouse/features/initialize/infra/nats"
	"github.com/rgfaber/go-vesca/greenhouse/initialize"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
	"github.com/rgfaber/go-vesca/sdk/infra/nats"
	sdk_nats_config "github.com/rgfaber/go-vesca/sdk/infra/nats/config"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	res := dig.New()
	res.Provide(config.NewConfig)
	res.Provide(sdk_nats_config.NewNatsConfig)
	res.Provide(mediator.SingletonDECBus)
	res.Provide(store.NewStore)
	res.Provide(initialize2.Feature)
	return res
}

var (
	Container          = dig.New()
	Config             = config.NewConfig()
	NatsConfig         = sdk_nats_config.NewNatsConfig()
	MemStore           = store.NewStore()
	MemBus             = mediator.SingletonDECBus()
	NatsBus            = nats.NewNatsBus(NatsConfig)
	InitializedEmitter = nats2.NewEmitter(MemBus, NatsBus, initialize.FACT_TOPIC, InitializedSerializer)
	Features           = []sdk.IFeature{
		initialize.NewFeature(MemBus,
			MemStore,
			InitializedEmitter),
		//kill.NewFeature(Root, MemBus),
		//measure.NewFeature(MemBus, MemStore),
	}
)
