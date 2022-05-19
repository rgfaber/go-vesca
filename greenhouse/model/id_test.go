package model

import (
	testing2 "github.com/rgfaber/go-vesca/go-scream/core/test"
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTHSensorTestID(t *testing.T) {
	id := NewGreenhouseTestID()
	assert.NotNil(t, id)
	assert.Equal(t, id.Prefix, config.GO_VESCA_GREENHOUSE_PREFIX)
	assert.Equal(t, id.Value, testing2.CLEAN_TEST_UUID)
}
