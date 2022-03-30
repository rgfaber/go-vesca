package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Feature struct {
	state    *model.Root
	commands Commands
	events   Events
}

func NewFeature(r *model.Root,
	chC Commands,
	chE Events,
) *Feature {
	return &Feature{
		commands: chC,
		events:   chE,
		state:    r,
	}
}

func (f *Feature) Respond(cmd Cmd) {
}

func (f *Feature) Run() {
	fmt.Println("measure.Feature is Up!")
}