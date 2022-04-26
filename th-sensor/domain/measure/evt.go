package measure

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

const EVT_TOPIC = "th_sensor:measured"

type Evt struct {
	sensorId    sdk.Identity
	neasurement model.Measurement
}

func NewEvt(sensorId sdk.Identity, m model.Measurement) *Evt {
	return &Evt{
		sensorId:    sensorId,
		neasurement: m,
	}
}
