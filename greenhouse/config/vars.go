package config

import (
	"github.com/rgfaber/go-vesca/go-scream/infra/nats"
	"github.com/rgfaber/go-vesca/go-scream/infra/nats/config"
)

var (
	natsCfg = config.NewNatsConfig()
	natsBus = nats.SingletonNatsBus(natsCfg)
)
