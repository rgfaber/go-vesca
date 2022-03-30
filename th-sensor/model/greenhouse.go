package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
)

type Greenhouse struct {
	ID      sdk.Identity `json:"id"`
	Details Details      `json:"details"`
}

func NewGreenhouse(id string) *Greenhouse {
	return &Greenhouse{
		ID: *sdk.NewIdentityFrom(config.GO_VESCA_GREENHOUSE_PREFIX, id),
	}
}
