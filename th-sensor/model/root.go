package model

import "github.com/rgfaber/go-vesca/sdk"

type Root struct {
	Id sdk.Identity `json:"id" bson:"id"`
}

func NewRoot(sensorId string) *Root {
	return &Root{
		Id: *sdk.NewIdentityFrom(TH_SENSOR, sensorId),
	}
}
