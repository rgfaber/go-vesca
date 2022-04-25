package model

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRoot(t *testing.T) {
	cfg := config.NewConfig()
	r := NewRoot(cfg.SensorId(), cfg.SensorName(), cfg.GreenhouseId())
	id := NewTHSensorId(cfg.SensorId())
	assert.NotNil(t, r)
	assert.NotNil(t, r.SensorId)
	assert.Equal(t, id.Id(), r.SensorId.Id())
}
