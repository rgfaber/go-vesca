package infra

import (
	config2 "github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBus(t *testing.T) {
	// Given
	cfg := config2.NewConfig()
	b := NewNatsBus(cfg)
	assert.NotNil(t, b, "Could not create initialize.NatsBus")
}
