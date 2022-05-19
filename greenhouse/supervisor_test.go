package main

import (
	testing2 "github.com/rgfaber/go-vesca/go-scream/core/test"
	"github.com/rgfaber/go-vesca/greenhouse/config/envars"
	"github.com/rgfaber/go-vesca/greenhouse/initialize"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewSupervisor(t *testing.T) {
	err := os.Setenv(envars.GO_VESCA_TH_SENSOR_ID, testing2.TEST_UUID)
	assert.Nil(t, err, "Error setting Environment Variable [%+v]", envars.GO_VESCA_TH_SENSOR_ID)
	supervisor := SingletonSupervisor(initialize.Feature)
	assert.NotNil(t, supervisor)
}
