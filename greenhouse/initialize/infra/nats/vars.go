package nats

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/contract"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/domain"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/topics"
	"github.com/rgfaber/go-vesca/sdk/dec"
	sdk_nats "github.com/rgfaber/go-vesca/sdk/infra/nats"
	"github.com/rgfaber/go-vesca/sdk/infra/nats/config"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"time"
)

var (
	cfg = config.NewNatsConfig()
	Bus = sdk_nats.TransientNatsBus(cfg)

	FactsChan = make(chan []byte)

	Emit2Nats = func(evt interfaces.IEvt) {
		e := evt.(*domain.Evt)
		f := domain.Evt2Fact(e)
		Bus.Publish(topics.NATS_FACT_TOPIC, contract.Fact2Json(f))
	}

	RequestByNats = func(hope interfaces.IHope) interfaces.IFbk {
		data := Bus.Request(topics.NATS_HOPE_TOPIC, contract.Hope2Json(hope), 5*60*time.Second)
		return domain.Json2Fbk(data)
	}

	OwnFactHandler = func(fact interfaces.IFact) {
		fmt.Printf("Received fact [%+v]", fact)
	}
	OwnHandler  = dec.NewFactHandler(OwnFactHandler, contract.Json2Fact)
	OwnListener = sdk_nats.NewNatsListener(Bus, topics.NATS_FACT_TOPIC, OwnHandler)
)
