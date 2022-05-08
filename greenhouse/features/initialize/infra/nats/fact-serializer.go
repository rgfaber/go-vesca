package nats

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/contract"
	"github.com/rgfaber/go-vesca/sdk"
)

type FactSerializer struct{}

func (ser FactSerializer) Serialize(fact sdk.IFact) ([]byte, error) {
	f := fact.(*contract.Fact)
	return json.Marshal(f)
}

func (ser FactSerializer) Deserialize(s []byte) sdk.IFact {
	var ret contract.Fact
	json.Unmarshal(s, &ret)
	return &ret
}

func NewFactSerializer() FactSerializer {
	return FactSerializer{}
}
