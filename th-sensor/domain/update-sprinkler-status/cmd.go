package update_sprinkler_status

import sprinkler_status "github.com/rgfaber/go-vesca/th-sensor/model/sprinkler-status"

type Cmd struct {
	sensorId  string
	newStatus sprinkler_status.SprinklerStatus
}

func NewCmd(sensorId string)
