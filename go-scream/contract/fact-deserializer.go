package contract

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type FactDeserializer struct {
	Handler func([]byte) interfaces.IFact
}

func (f *FactDeserializer) Deserialize(data []byte) interfaces.IFact {
	return f.Handler(data)
}

func NewFactDeserializer(handler func([]byte) interfaces.IFact) *FactDeserializer {
	return &FactDeserializer{
		Handler: handler,
	}
}
