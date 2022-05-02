package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
)

func NewGreenhouseTestID() *sdk.Identity {
	return sdk.NewIdentityFrom(config.GO_VESCA_GREENHOUSE_PREFIX, sdk.TEST_UUID)
}

func NewGreenhouseID(id string) *sdk.Identity {
	return sdk.NewIdentityFrom(config.GO_VESCA_GREENHOUSE_PREFIX, id)
}
