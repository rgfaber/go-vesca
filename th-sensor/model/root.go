package model

import "github.com/rgfaber/go-vesca/sdk"

type Root struct {
	Id sdk.Identity `json:"id" bson:"id"`
}

func NewRoot() *Root {
	id := sdk.NewIdentity("th_sensor")
	return &Root{Id: *id}
}
