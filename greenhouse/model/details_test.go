package model

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDetails(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	// When
	d := NewDetails(cfg.GreenhouseName(), "")
	// Then
	assert.NotNil(t, d)
	assert.Equal(t, cfg.GreenhouseName(), d.Name)
}
