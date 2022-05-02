package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
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
