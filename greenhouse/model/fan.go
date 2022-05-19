package model

import (
	"github.com/rgfaber/go-vesca/go-scream/core"
	"github.com/rgfaber/go-vesca/greenhouse/config"
)

type Fan struct {
	ID      *core.Identity `json:"id"`
	Details *Details       `json:"details"`
}

func NewFan(id string, name string) *Fan {
	return &Fan{
		ID:      core.NewIdentityFrom(config.GO_VESCA_FAN_PREFIX, id),
		Details: NewDetails(name, ""),
	}
}
