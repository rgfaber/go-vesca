package contract

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"log"
)

var (
	Hope2Json = func(hope interfaces.IHope) []byte {
		h := hope.(*Hope)
		res, e := json.Marshal(h)
		if e != nil {
			return nil
		}
		return res
	}
	Json2Hope = func(data []byte) interfaces.IHope {
		var hp Hope
		json.Unmarshal(data, &hp)
		return &hp
	}
	Fact2Json = func(fact interfaces.IFact) []byte {
		f := fact.(*Fact)
		res, e := json.Marshal(f)
		if e != nil {
			log.Fatal(e)
			return nil
		}
		return res
	}

	Json2Fact = func(data []byte) interfaces.IFact {
		var f Fact
		err := json.Unmarshal(data, &f)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return &f
	}
)
