package dec

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type HopeHandler struct {
	Str2Hope   func([]byte) interfaces.IHope
	Hope2Cmd   func(hope interfaces.IHope) interfaces.ICmd
	Fbk2String func(fbk interfaces.IFbk) []byte
	Aggregate  interfaces.IAggregate
}

func NewHopeHandler(aggregate interfaces.IAggregate,
	hope2Cmd func(hope interfaces.IHope) interfaces.ICmd,
	str2Hope func([]byte) interfaces.IHope,
	fbk2Str func(fbk interfaces.IFbk) []byte) *HopeHandler {
	return &HopeHandler{
		Aggregate:  aggregate,
		Str2Hope:   str2Hope,
		Hope2Cmd:   hope2Cmd,
		Fbk2String: fbk2Str,
	}
}

func (h *HopeHandler) Handle(data []byte) []byte {
	hp := h.Str2Hope(data)
	cmd := h.Hope2Cmd(hp)
	fbk := h.Aggregate.Attempt(cmd)
	return h.Fbk2String(fbk)
}
