package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmitter(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	nats := infra.NewNatsBus(cfg)
	// When
	emitter := NewEmitter(nats)
	// Then
	assert.NotNil(t, emitter)
}

func TestIfEmitterImplementsIEmitter(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	bus := infra.NewNatsBus(cfg)
	e := NewEmitter(bus)
	// When
	ok := dec.ImplementsIEmitter(e)
	//
	assert.True(t, ok)
}
