package dec

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"log"
)

type HopeHandler struct {
	handler          func(hope interfaces.IHope) (interfaces.IFbk, error)
	hopeDeserializer interfaces.IHopeDeserializer
	fbkSerializer    interfaces.IFbkSerializer
}

func NewHopeHandler(handler func(hope interfaces.IHope) (interfaces.IFbk, error),
	hopeDeserializer interfaces.IHopeDeserializer,
	fbkSerializer interfaces.IFbkSerializer) *HopeHandler {
	return &HopeHandler{
		handler:          handler,
		hopeDeserializer: hopeDeserializer,
		fbkSerializer:    fbkSerializer,
	}
}

func (h *HopeHandler) Handle(data []byte) []byte {
	hp := h.hopeDeserializer.Deserialize(data)
	fbk, errHandle := h.handler(hp)
	if errHandle != nil {
		log.Fatal(errHandle)
		return nil
	}
	res, errSer := h.fbkSerializer.Serialize(fbk)
	if errSer != nil {
		log.Fatal(errSer)
		return nil
	}
	return res
}
