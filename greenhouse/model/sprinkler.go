package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
)

type Sprinkler struct {
	ID      sdk.Identity
	Details Details
}

func NewSprinkler(id string, name string) *Sprinkler {
	return &Sprinkler{
		ID:      *sdk.NewIdentityFrom(config.GO_VESCA_SPRINKLER_PREFIX, id),
		Details: *NewDetails(name, ""),
	}
}
