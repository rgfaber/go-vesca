package infra

import (
	"github.com/rgfaber/go-vesca/sdk/domain"
)

type Mediator chan domain.ITopic

func NewMediator(size int) Mediator {
	return make(Mediator, size)
}
