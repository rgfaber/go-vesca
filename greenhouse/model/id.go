package model

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/core/mocks"
)

func NewGreenhouseTestID() *core.Identity {
	return core.NewIdentityFrom(config.GO_VESCA_GREENHOUSE_PREFIX, mocks.TEST_UUID)
}

func NewGreenhouseID(id string) *core.Identity {
	return core.NewIdentityFrom(config.GO_VESCA_GREENHOUSE_PREFIX, id)
}
