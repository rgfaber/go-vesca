package domain

import (
	"encoding/json"
	contract2 "github.com/rgfaber/go-vesca/greenhouse/features/replace-fan/contract"
	"github.com/rgfaber/go-vesca/greenhouse/features/replace-fan/infra/memory"
	"github.com/rgfaber/go-vesca/sdk/contract"
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"log"
)

var (
	Root = NewAggregate(memory.Store, memory.Mediator)

	Hope2Cmd = func(hope contract.IHope) domain.ICmd {
		if hope == nil {
			return nil
		}
		h := hope.(*contract2.Hope)
		id := core.NewIdentityFromAggregateId(hope.AggregateId())
		traceId := hope.TraceId()
		return NewCmd(id, traceId, h.Fan)
	}

	Evt2Fact = func(evt domain.IEvt) domain.IFact {
		e := evt.(*Evt)
		return contract2.NewFact(e.AggregateId().Id(), e.TraceId(), e.Fan)
	}

	Fbk2Json = func(fbk domain.IFbk) []byte {
		f := fbk.(*Fbk)
		res, err := json.Marshal(f)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return res
	}

	Json2Fbk = func(data []byte) domain.IFbk {
		var fbk Fbk
		err := json.Unmarshal(data, &fbk)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return &fbk
	}
)
