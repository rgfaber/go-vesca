package measure

import "github.com/rgfaber/go-vesca/greenhouse/model"

type Fact struct {
	sensorId    string
	measurement model.Settings
}

func NewFact() *Fact {
	return &Fact{}
}
