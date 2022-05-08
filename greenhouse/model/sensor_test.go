package model

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSensor(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	id := cfg.SensorId()
	// When
	sensor := NewSensor(id, cfg.SensorName())
	// Then
	assert.NotNil(t, sensor)
	assert.Equal(t, sdk.CLEAN_TEST_UUID, sensor.ID.Value)
}
