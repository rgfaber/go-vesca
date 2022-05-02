package model

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/sdk"
)

type Fan struct {
	ID      *sdk.Identity `json:"id"`
	Details *Details      `json:"details"`
}

func NewFan(id string, name string) *Fan {
	return &Fan{
		ID:      sdk.NewIdentityFrom(config.GO_VESCA_FAN_PREFIX, id),
		Details: NewDetails(name, ""),
	}
}
