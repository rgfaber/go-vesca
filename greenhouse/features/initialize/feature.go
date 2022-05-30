package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/infra/nats"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/features"
	"github.com/rgfaber/go-vesca/sdk/infra"
)

var (
	listeners  = infra.BuildListeners(nats.OwnListener)
	responders = infra.BuildResponders(nats.Responder)
	handlers   = domain.BuildHandlers(nats.Emitter)
	Feature    = features.NewFeature(responders, listeners, handlers)
)
