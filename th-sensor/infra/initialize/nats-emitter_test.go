package initialize

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
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
