package dec

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type Mediator chan interfaces.ITopic

func NewMediator(size int) Mediator {
	return make(Mediator, size)
}
