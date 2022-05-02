package model

import (
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
	assert.Equal(t, id, sensor.ID.Value)
}
