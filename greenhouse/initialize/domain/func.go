package domain

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/go-scream/core"
	"github.com/rgfaber/go-vesca/go-scream/interfaces"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/contract"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/infra/memory"
	"log"
)

var (
	Root = NewAggregate(memory.Store, memory.Mediator)

	Hope2Cmd = func(hope interfaces.IHope) interfaces.ICmd {
		if hope == nil {
			return nil
		}
		h := hope.(*contract.Hope)
		id := core.NewIdentityFromAggregateId(hope.AggregateId())
		traceId := hope.TraceId()
		return NewCmd(id, traceId, h.Details, h.Settings, h.Sensor, h.Fan, h.Sprinkler)
	}

	Evt2Fact = func(evt interfaces.IEvt) interfaces.IFact {
		e := evt.(*Evt)
		return contract.NewFact(e.AggregateId().Id(), e.TraceId(), e.Greenhouse)
	}

	Fbk2Json = func(fbk interfaces.IFbk) []byte {
		f := fbk.(*Fbk)
		res, err := json.Marshal(f)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return res
	}

	Json2Fbk = func(data []byte) interfaces.IFbk {
		var fbk Fbk
		err := json.Unmarshal(data, &fbk)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return &fbk
	}
)
