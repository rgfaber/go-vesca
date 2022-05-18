package dec

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"sync"
)

type FactHandler struct {
	handler      func(fact interfaces.IFact)
	deserializer func(data []byte) interfaces.IFact
}

func NewFactHandler(handler func(fact interfaces.IFact),
	deserializer func(data []byte) interfaces.IFact) *FactHandler {
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
