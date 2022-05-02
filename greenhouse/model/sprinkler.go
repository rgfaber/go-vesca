package model

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/sdk"
)

type Sprinkler struct {
	ID      *sdk.Identity `json:"id"`
	Details *Details      `json:"details"`
}

func NewSprinkler(id string, name string) *Sprinkler {
	return &Sprinkler{
		ID:      sdk.NewIdentityFrom(config.GO_VESCA_SPRINKLER_PREFIX, id),
		Details: NewDetails(name, ""),
	}
}
