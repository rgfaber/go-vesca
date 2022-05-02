package main

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/envars"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewSupervisor(t *testing.T) {
	err := os.Setenv(envars.GO_VESCA_TH_SENSOR_ID, sdk.TEST_UUID)
	assert.Nil(t, err, "Error setting Environment Variable [%+v]", envars.GO_VESCA_TH_SENSOR_ID)
	cfg := config.NewConfig()
	supervisor := NewSupervisor(cfg, MemBus, Features)
	assert.NotNil(t, supervisor)
	assert.NotNil(t, supervisor.bus)
}
