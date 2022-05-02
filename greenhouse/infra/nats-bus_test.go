package infra

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBus(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	// When
	b := NewNatsBus(cfg)
	// Then
	assert.NotNil(t, b, "Could not create initialize.NatsBus")
}

func TestNatsBusImplementsINatsBus(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	b := NewNatsBus(cfg)
	// When
	ok := ImplementsINatsBus(b)
	// Then
	assert.True(t, ok)
}
