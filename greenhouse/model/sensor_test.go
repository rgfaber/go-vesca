package model

import (
	testing2 "github.com/rgfaber/go-vesca/go-scream/core/test"
	"github.com/rgfaber/go-vesca/greenhouse/config"
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
	assert.Equal(t, testing2.CLEAN_TEST_UUID, sensor.ID.Value)
}
