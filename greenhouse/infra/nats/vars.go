package nats

import (
	"github.com/rgfaber/go-vesca/sdk/infra/nats"
	"github.com/rgfaber/go-vesca/sdk/infra/nats/config"
)

var (
	natsCfg = config.NewNatsConfig()

	NatsBus = nats.SingletonNatsBus(natsCfg)
)
