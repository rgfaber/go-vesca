package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/initialize/contract"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/domain"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/infra/memory"
	infra_nats "github.com/rgfaber/go-vesca/greenhouse/initialize/infra/nats"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/topics"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/sdk/infra/nats"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

var (
	hopeHandler   = dec.NewHopeHandler(domain.Root, domain.Hope2Cmd, contract.Json2Hope, domain.Fbk2Json)
	natsResponder = nats.NewNatsResponder(infra_nats.Bus, topics.NATS_HOPE_TOPIC, hopeHandler)
	natsEmitter   = dec.NewEventHandler(memory.Mediator, topics.EVT_TOPIC, infra_nats.Emit2Nats)

	listeners  = interfaces.BuildListeners()
	responders = interfaces.BuildResponders(natsResponder)
	handlers   = interfaces.BuildHandlers(natsEmitter)

	Feature = dec.NewFeature(responders, listeners, handlers)
)
