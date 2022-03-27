package model

import "../sdk"

type Root struct {
	id sdk.Identity `json:"id" bson:"id"`
}
