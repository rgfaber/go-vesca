package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
)

func NewTHSensorTestID() *sdk.Identity {
	return sdk.NewIdentityFrom(config.GO_VESCA_TH_SENSOR_PREFIX, sdk.TEST_UUID)
}

func NewTHSensorId(cfg *config.Config) *sdk.Identity {
	return sdk.NewIdentityFrom(config.GO_VESCA_TH_SENSOR_PREFIX, cfg.SensorId())
}
