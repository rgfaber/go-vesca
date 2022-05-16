package measure

import "github.com/rgfaber/go-vesca/greenhouse/model"

type Fact struct {
	AggregateId string          `json:"aggregate_id"`
	Measurement *model.Settings `json:"measurement"`
}

func NewFact(aggregateId string, measurement *model.Settings) *Fact {
	return &Fact{
		AggregateId: aggregateId,
		Measurement: measurement,
	}
}
