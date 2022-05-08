package dec

import "github.com/rgfaber/go-vesca/sdk/interfaces"

type HopeDeserializer struct {
	handler func([]byte) interfaces.IHope
}

func NewHopeDeserializer(handler func([]byte) interfaces.IHope) *HopeDeserializer {
	return &HopeDeserializer{
		handler: handler,
	}
}

func (d *HopeDeserializer) Deserialize(data []byte) interfaces.IHope {
	return d.handler(data)
}
