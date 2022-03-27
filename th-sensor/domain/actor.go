package domain

import "github.com/rgfaber/go-vesca/sdk"

type Actor struct {
	Id *sdk.Identity
}

func NewActor(sensorId string) *Actor {
	id := sdk.NewIdentityFrom("th_sensor", sensorId)
	return &Actor{Id: id}
}
