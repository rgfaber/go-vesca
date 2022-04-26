package update_fan_status

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	fan_status "github.com/rgfaber/go-vesca/th-sensor/model/fan-status"
)

type Cmd struct {
	sensorId  string
	newStatus fan_status.FanStatus
}

func NewCmd(sensorId string, newStatus fan_status.FanStatus) *Cmd {
	return &Cmd{
		sensorId:  sensorId,
		newStatus: newStatus,
	}
}

func (c *Cmd) AggregateId() sdk.IIdentity {
	return model.NewTHSensorId(c.sensorId)
}
