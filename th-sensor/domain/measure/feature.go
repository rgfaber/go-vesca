package measure

import (
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Feature struct {
	state *model.Root
	chCmd ChCmd
	chEvt ChEvt
}

func NewFeature(state *model.Root,
	chCmd ChCmd,
	chEvt ChEvt) *Feature {
	return &Feature{
		state: state,
		chCmd: chCmd,
		chEvt: chEvt,
	}
}

func (f *Feature) listen() {

}

func (f *Feature) Run() {
	go f.listen()

}
