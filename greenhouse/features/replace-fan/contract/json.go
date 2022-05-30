package contract

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/sdk/contract"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"log"
)

var (
	Hope2Json = func(hope contract.IHope) []byte {
		h := hope.(*Hope)
		res, err := json.Marshal(h)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return res
	}
	Json2Hope = func(data []byte) contract.IHope {
		var hp Hope
		err := json.Unmarshal(data, &hp)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return &hp
	}
	Fact2Json = func(fact domain.IFact) []byte {
		f := fact.(*Fact)
		res, err := json.Marshal(f)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return res
	}

	Json2Fact = func(data []byte) domain.IFact {
		var f Fact
		err := json.Unmarshal(data, &f)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return &f
	}
)
