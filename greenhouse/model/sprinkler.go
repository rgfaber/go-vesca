package model

import (
	"github.com/rgfaber/go-vesca/go-scream/core"
	"github.com/rgfaber/go-vesca/greenhouse/config"
)

type Sprinkler struct {
	ID      *core.Identity `json:"id"`
	Details *Details       `json:"details"`
}

func NewSprinkler(id string, name string) *Sprinkler {
	return &Sprinkler{
		ID:      core.NewIdentityFrom(config.GO_VESCA_SPRINKLER_PREFIX, id),
		Details: NewDetails(name, ""),
	}
}
