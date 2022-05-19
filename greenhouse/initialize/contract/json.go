package contract

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/go-scream/interfaces"
	"log"
)

var (
	Hope2Json = func(hope interfaces.IHope) []byte {
		h := hope.(*Hope)
		res, err := json.Marshal(h)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return res
	}
	Json2Hope = func(data []byte) interfaces.IHope {
		var hp Hope
		err := json.Unmarshal(data, &hp)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		return &hp
	}
	Fact2Json = func(fact interfaces.IFact) []byte {
		f := fact.(*Fact)
		res, err := json.Marshal(f)
		if err != nil {
			log.Fatal(err)
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
