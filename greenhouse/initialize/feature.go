package initialize

import (
	"github.com/rgfaber/go-vesca/go-scream/dec"
	"github.com/rgfaber/go-vesca/go-scream/interfaces"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/infra/nats"
)

var (
	listeners  = interfaces.BuildListeners(nats.OwnListener)
	responders = interfaces.BuildResponders(nats.Responder)
	handlers   = interfaces.BuildHandlers(nats.Emitter)
	Feature    = dec.NewFeature(responders, listeners, handlers)
)
