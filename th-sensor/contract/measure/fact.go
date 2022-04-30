package measure

import "github.com/rgfaber/go-vesca/th-sensor/model"

type Fact struct {
	sensorId    string
	measurement model.Measurement
}

func NewFact() *Fact {
	return &Fact{}
}
