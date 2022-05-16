package model

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTHSensorTestID(t *testing.T) {
	id := NewGreenhouseTestID()
	assert.NotNil(t, id)
	assert.Equal(t, id.Prefix, config.GO_VESCA_GREENHOUSE_PREFIX)
	assert.Equal(t, id.Value, testing2.CLEAN_TEST_UUID)
}
