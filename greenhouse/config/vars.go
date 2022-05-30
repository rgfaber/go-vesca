package config

import (
	"github.com/rgfaber/go-vesca/sdk/infra/nats"
	"github.com/rgfaber/go-vesca/sdk/infra/nats/config"
)

var (
	natsCfg = config.NewNatsConfig()
	natsBus = nats.SingletonNatsBus(natsCfg)
)
