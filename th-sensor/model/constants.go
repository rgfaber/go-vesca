package model

import "os"

const (
	TH_SENSOR    = "th_sensor"
	TH_SENSOR_ID = "TH_SENSOR_ID"
)

var (
	SensorId = os.Getenv(TH_SENSOR_ID)
)
