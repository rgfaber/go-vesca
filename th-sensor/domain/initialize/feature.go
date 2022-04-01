package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Feature struct {
	state    *model.Root
	mediator sdk.Mediator
}

func NewFeature(r *model.Root,
	mediator sdk.Mediator,
) *Feature {
	return &Feature{
		mediator: mediator,
		state:    r,
	}
}

func (f *Feature) Respond(cmd Cmd) {
}

func (f *Feature) Run() {
	fmt.Println("measure.Feature is Up!")
}
