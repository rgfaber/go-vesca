package infra

import (
	config2 "github.com/rgfaber/go-vesca/th-sensor/config"
)

var (
	config = config2.NewConfig()
	bus    = NewBus(config)
	//	NatsEmitter = NewEmitter(nats.Connect(config.NATSUrl))
)
