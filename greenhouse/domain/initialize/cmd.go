package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
)

const CMD_TOPIC = "th_sensor:initialize"

type Cmd struct {
	aggregateId sdk.Identity
	traceId     string
	name        string
	settings    model.Settings
	sensor      model.Sensor
	fan         model.Fan
	sprinkler   model.Sprinkler
}

func NewCmd(id sdk.Identity, traceId string, name string, settings model.Settings, sensor model.Sensor, fan model.Fan, sprinkler model.Sprinkler) *Cmd {
	return &Cmd{
		aggregateId: id,
		name:        name,
		traceId:     traceId,
		settings:    settings,
		sensor:      sensor,
		fan:         fan,
		sprinkler:   sprinkler,
	}
}

func (c *Cmd) AggregateId() sdk.IIdentity {
	return &c.aggregateId
}
