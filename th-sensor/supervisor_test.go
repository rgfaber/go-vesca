package main

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/envars"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestNewSupervisor(t *testing.T) {
	err := os.Setenv(envars.GO_VESCA_TH_SENSOR_ID, sdk.TEST_UUID)
	assert.Nil(t, err, "Error setting Environment Variable [%+v]", envars.GO_VESCA_TH_SENSOR_ID)
	cfg := config.NewConfig()
	supervisor := NewSupervisor(cfg, nil)
	assert.NotNil(t, supervisor)
	assert.NotNil(t, supervisor.State)
	assert.NotNil(t, supervisor.State.ID)
	assert.Equal(t, config.GO_VESCA_TH_SENSOR_PREFIX, supervisor.State.ID.Prefix)
	ts := strings.Replace(cfg.SensorId(), "-", "", -1)
	assert.Equal(t, ts, supervisor.State.ID.Value)
}
