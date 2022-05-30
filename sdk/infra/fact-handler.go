package infra

import (
	"github.com/rgfaber/go-vesca/sdk/domain"
	"sync"
)

type IFactHandler interface {
	Handle(data []byte)
}

func ImplementsIFactHandler(handler IFactHandler) bool {
	return true
}

type FactHandler struct {
	handler      func(fact domain.IFact)
	deserializer func(data []byte) domain.IFact
}

func NewFactHandler(handler func(fact domain.IFact),
	deserializer func(data []byte) domain.IFact) *FactHandler {
	return &FactHandler{
		deserializer: deserializer,
		handler:      handler,
	}
}

var lock = sync.Mutex{}

func (fh *FactHandler) Handle(data []byte) {
	if data != nil {
		lock.Lock()
		defer lock.Unlock()
		fact := fh.deserializer(data)
		fh.handler(fact)
	}
	return
}
