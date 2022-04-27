package update_fan_status

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	fan_status "github.com/rgfaber/go-vesca/th-sensor/model/fan-status"
)

const EVT_TOPIC = "th-sensor:fan-status-updated"

type Evt struct {
	sensorId  string
	newStatus fan_status.FanStatus
}

func (e *Evt) AggregateId() sdk.IIdentity {
	return model.NewTHSensorId(e.sensorId)
}

func NewEvt(sensorId string, newStatus fan_status.FanStatus) *Evt {
	return &Evt{
		sensorId:  sensorId,
		newStatus: newStatus,
	}
}
