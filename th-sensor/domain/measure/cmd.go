package measure

import "github.com/rgfaber/go-vesca/th-sensor/model"

const CMD_TOPIC = "th_sensor:measure"

type Cmd struct {
	Measurement *model.Measurement
}

func NewCmd(m *model.Measurement) *Cmd {
	return &Cmd{
		Measurement: m,
	}
}
