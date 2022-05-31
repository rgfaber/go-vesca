package contract

import (
	"encoding/json"
	domain2 "github.com/rgfaber/go-vesca/greenhouse/features/initialize/domain"
	"github.com/rgfaber/go-vesca/sdk/contract"
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"log"
)

var (
	Hope2Cmd = func(hope contract.IHope) domain.ICmd {
		if hope == nil {
			return nil
		}
		h := hope.(*Hope)
		id := core.NewIdentityFromAggregateId(hope.AggregateId())
		traceId := hope.TraceId()
		return domain2.NewCmd(id, traceId, h.Details, h.Settings, h.Sensor, h.Fan, h.Sprinkler)
	}

	Evt2Fact = func(evt domain.IEvt) domain.IFact {
		e := evt.(*domain2.Evt)
		return NewFact(e.AggregateId().Id(), e.TraceId(), e.Greenhouse)
	}

	Fbk2Json = func(fbk domain.IFbk) []byte {
		f := fbk.(*domain2.Fbk)
		res, err := json.Marshal(f)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return res
	}

	Json2Fbk = func(data []byte) domain.IFbk {
		var fbk domain2.Fbk
		err := json.Unmarshal(data, &fbk)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return &fbk
	}
)
