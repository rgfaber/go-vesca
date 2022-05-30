package model

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/sdk/core"
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
