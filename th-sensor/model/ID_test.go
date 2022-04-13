package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewTHSensorTestID() *sdk.Identity {
	return sdk.NewIdentityFrom(config.GO_VESCA_TH_SENSOR_PREFIX, sdk.TEST_UUID)
}

func TestNewTHSensorTestID(t *testing.T) {
	id := NewTHSensorTestID()
	assert.NotNil(t, id)

}
