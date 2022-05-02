package infra

import "github.com/rgfaber/go-vesca/th-sensor/config"

var (
	Nats = NewNatsBus(config.NewConfig())
	//	NatsEmitter = NewEmitter(nats.Connect(config.NATSUrl))
)
