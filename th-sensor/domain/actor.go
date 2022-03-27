package domain

import "github.com/rgfaber/go-vesca/sdk"

type Actor struct {
	Id *sdk.Identity
}

func NewActor(sensorId string) *Actor {
	return &Actor{Id: sdk.NewIdentityFrom("th_sensor", sensorId)}
}
