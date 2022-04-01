package measure

import "github.com/rgfaber/go-vesca/th-sensor/model"

const EVT_TOPIC = "th_sensor:measured"

type Evt struct {
	Measurement *model.Measurement
}

func NewEvt(m *model.Measurement) *Evt {
	return &Evt{
		Measurement: m,
	}
}
