package nats

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/contract"
	domain2 "github.com/rgfaber/go-vesca/greenhouse/features/initialize/domain"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/infra/memory"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/topics"
	contract2 "github.com/rgfaber/go-vesca/sdk/contract"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra"
	sdk_nats "github.com/rgfaber/go-vesca/sdk/infra/nats"
	"github.com/rgfaber/go-vesca/sdk/infra/nats/config"
	"time"
)

var (
	cfg = config.NewNatsConfig()
	Bus = sdk_nats.TransientNatsBus(cfg)

	FactsChan = make(chan []byte)

	Emit2Nats = func(evt domain.IEvt) {
		e := evt.(*domain2.Evt)
		f := domain2.Evt2Fact(e)
		Bus.Publish(topics.NATS_FACT_TOPIC, contract.Fact2Json(f))
	}

	RequestByNats = func(hope contract2.IHope) domain.IFbk {
		data := Bus.Request(topics.NATS_HOPE_TOPIC, contract.Hope2Json(hope), 5*60*time.Second)
		return domain2.Json2Fbk(data)
	}

	hopeHandler = infra.NewHopeHandler(domain2.Root, domain2.Hope2Cmd, contract.Json2Hope, domain2.Fbk2Json)

	Responder = sdk_nats.NewNatsResponder(Bus, topics.NATS_HOPE_TOPIC, hopeHandler)

	Emitter = domain.NewEventHandler(memory.Mediator, topics.EVT_TOPIC, Emit2Nats)

	OwnFactHandler = func(fact domain.IFact) {
		fmt.Printf("Received fact [%+v]", fact)
	}
	OwnHandler  = infra.NewFactHandler(OwnFactHandler, contract.Json2Fact)
	OwnListener = sdk_nats.NewNatsListener(Bus, topics.NATS_FACT_TOPIC, OwnHandler)
)
