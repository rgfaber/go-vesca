package contract

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type FactSerializer struct {
	Handler func(fact interfaces.IFact) ([]byte, error)
}

func NewFactSerializer(handler func(fact interfaces.IFact) ([]byte, error)) *FactSerializer {
	return &FactSerializer{
		Handler: handler,
	}
}

func (s *FactSerializer) Serialize(fact interfaces.IFact) ([]byte, error) {
	return s.Handler(fact)
}
