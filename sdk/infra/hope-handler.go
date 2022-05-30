package infra

import (
	"github.com/rgfaber/go-vesca/sdk/contract"
	"github.com/rgfaber/go-vesca/sdk/domain"
)

type IHopeHandler interface {
	Handle(hope []byte) []byte
}

func ImplementsIHopeHandler(handler IHopeHandler) bool {
	return true
}

type HopeHandler struct {
	Str2Hope   func([]byte) contract.IHope
	Hope2Cmd   func(hope contract.IHope) domain.ICmd
	Fbk2String func(fbk domain.IFbk) []byte
	Aggregate  domain.IAggregate
}

func NewHopeHandler(aggregate domain.IAggregate,
	hope2Cmd func(hope contract.IHope) domain.ICmd,
	str2Hope func([]byte) contract.IHope,
	fbk2Str func(fbk domain.IFbk) []byte) *HopeHandler {
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
