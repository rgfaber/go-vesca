package model

import (
	"github.com/rgfaber/go-vesca/go-scream/core"
	"github.com/rgfaber/go-vesca/go-scream/core/test"
	"github.com/rgfaber/go-vesca/greenhouse/config"
)

func NewGreenhouseTestID() *core.Identity {
	return core.NewIdentityFrom(config.GO_VESCA_GREENHOUSE_PREFIX, test.TEST_UUID)
}

func NewGreenhouseID(id string) *core.Identity {
	return core.NewIdentityFrom(config.GO_VESCA_GREENHOUSE_PREFIX, id)
}
