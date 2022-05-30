package nats

import (
	"github.com/rgfaber/go-vesca/sdk/infra/nats/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBus(t *testing.T) {
	// Given
	cfg := config.NewNatsConfig()
	// When
	b := SingletonNatsBus(cfg)
	// Then
	assert.NotNil(t, b, "Could not create initialize.NatsBus")
}

func TestNatsBusImplementsINatsBus(t *testing.T) {
	// Given
	cfg := config.NewNatsConfig()
	b := SingletonNatsBus(cfg)
	// When
	ok := ImplementsINatsBus(b)
	// Then
	assert.True(t, ok)
}
