package initialize

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
	"github.com/rgfaber/go-vesca/sdk/infra/nats"
	"github.com/rgfaber/go-vesca/sdk/infra/nats/config"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"log"
)

var (
	natsCfg  = config.NewNatsConfig()
	natsBus  = nats.NewNatsBus(natsCfg)
	Store    = store.NewStore()
	Mediator = mediator.SingletonDECBus()
	Root     = NewAggregate(Store, Mediator)
	hope2Cmd = func(hope interfaces.IHope) interfaces.ICmd {
		if hope == nil {
			return nil
		}
		h := hope.(*Hope)
		id := core.NewIdentityFromAggregateId(hope.AggregateId())
		traceId := hope.TraceId()
		return NewCmd(id, traceId, h.Details, h.Settings, h.Sensor, h.Fan, h.Sprinkler)
	}
	hope2Json = func(hope interfaces.IHope) []byte {
		h := hope.(*Hope)
		res, e := json.Marshal(h)
		if e != nil {
			return nil
		}
		return res
	}
	json2Hope = func(data []byte) interfaces.IHope {
		var hp Hope
		json.Unmarshal(data, &hp)
		return &hp
	}
	fbk2Json = func(fbk interfaces.IFbk) []byte {
		f := fbk.(*Fbk)
		res, err := json.Marshal(f)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return res
	}
	fact2Json = func(fact interfaces.IFact) []byte {
		f := fact.(*Fact)
		res, e := json.Marshal(f)
		if e == nil {
			log.Fatal(e)
			return nil
		}
		return res
	}
	evt2Fact = func(evt interfaces.IEvt) interfaces.IFact {
		e := evt.(*Evt)
		return NewFact(e.AggregateId().Id(), e.TraceId(), e.Greenhouse)
	}
	hopeHandler = dec.NewHopeHandler(Root, hope2Cmd, json2Hope, fbk2Json)

	natsResponder = nats.NewNatsResponder(natsBus, NATS_HOPE_TOPIC, hopeHandler)

	evt2Nats = func(evt interfaces.IEvt) {
		e := evt.(*Evt)
		f := evt2Fact(e)
		natsBus.Publish(NATS_FACT_TOPIC, fact2Json(f))
	}
	natsEmitter = dec.NewEventHandler(Mediator, EVT_TOPIC, evt2Nats)

	responders = interfaces.BuildResponders(natsResponder)
	listeners  = interfaces.BuildListeners()
	handlers   = interfaces.BuildHandlers(natsEmitter)

	Feature = dec.NewFeature(responders, listeners, handlers)
)
