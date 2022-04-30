package measure

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

const EVT_TOPIC = "th_sensor:measured"

type Evt struct {
	sensorId    string
	temperature float64
	humidity    float64
}

func (e Evt) AggregateId() sdk.IIdentity {
	return model.NewTHSensorId(e.sensorId)

}

func NewEvt(sensorId string, temperature float64, humidity float64) *Evt {
	return &Evt{
		sensorId:    sensorId,
		temperature: temperature,
		humidity:    humidity,
	}
}
