package dec

import "github.com/rgfaber/go-vesca/sdk/interfaces"

type FactHandler struct {
	handler      func(fact interfaces.IFact)
	deserializer interfaces.IFactDeserializer
}

func NewFactHandler(deserializer interfaces.IFactDeserializer, handler func(fact interfaces.IFact)) *FactHandler {
	return &FactHandler{
		deserializer: deserializer,
		handler:      handler,
	}
}

func (fh *FactHandler) Handle(data []byte) {
	if data != nil {
		fact := fh.deserializer.Deserialize(data)
		fh.handler(fact)
	}
	return
}
