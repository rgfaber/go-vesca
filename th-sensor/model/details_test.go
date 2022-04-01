package model

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDetails(t *testing.T) {
	cfg := config.NewConfig()
	d := NewDetails(cfg.GreenhouseName())
	assert.NotNil(t, d)
}
