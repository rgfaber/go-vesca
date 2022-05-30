package main

import (
	"github.com/rgfaber/go-vesca/greenhouse/config/envars"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/mocks"
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
