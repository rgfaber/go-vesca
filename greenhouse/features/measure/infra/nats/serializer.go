package nats

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/contract"
)

type Serializer struct{}

func (ser Serializer) Serialize(fact *contract.Fact) []byte {
	s, err := json.Marshal(*fact)
	if err != nil {
		panic(err)
	}
	return s
}

func (ser Serializer) Deserialize(s []byte) *contract.Fact {
	var ret contract.Fact
	json.Unmarshal(s, &ret)
	return &ret
}

func NewSerializer() Serializer {
	return Serializer{}
}
