package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTHSensorTestID(t *testing.T) {
	id := NewGreenhouseTestID()
	assert.NotNil(t, id)
	assert.Equal(t, id.Prefix, config.GO_VESCA_TH_SENSOR_PREFIX)
	assert.Equal(t, id.Value, sdk.CLEAN_TEST_UUID)
}
