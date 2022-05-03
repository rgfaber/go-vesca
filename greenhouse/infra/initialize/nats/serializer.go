package nats

import (
	"encoding/json"
	"github.com/rgfaber/go-vesca/greenhouse/infra/contract/initialize"
)

type Serializer struct{}

func (ser Serializer) Serialize(fact *initialize.Fact) []byte {
	s, err := json.Marshal(*fact)
	if err != nil {
		panic(err)
	}
	return s
}

func (ser Serializer) Deserialize(s []byte) *initialize.Fact {
	var ret initialize.Fact
	json.Unmarshal(s, &ret)
	return &ret
}

func NewSerializer() Serializer {
	return Serializer{}
}
