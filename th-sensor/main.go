package main

import (
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"os"
)

const (
	TH_SENSOR    = "th_sensor"
	TH_SENSOR_ID = "TH_SENSOR_ID"
)

func main() {
	sensorId = os.Getenv(TH_SENSOR_ID)
	act := domain.NewActor(sensorId)

}
