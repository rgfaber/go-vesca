package nats

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/nats"
	sdk_nats_config "github.com/rgfaber/go-vesca/sdk/nats/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListener(t *testing.T) {
	// Given
	cfg := sdk_nats_config.NewNatsConfig()
	memBus := sdk.NewDECBus()
	natsBus := nats.NewNatsBus(cfg)
	// When
	lst := NewListener(memBus, natsBus)
	// Then
	assert.NotNil(t, lst)
}

func TestListenerImplementsIListener(t *testing.T) {
	// Given
	cfg := sdk_nats_config.NewNatsConfig()
	memBus := sdk.NewDECBus()
	natsBus := nats.NewNatsBus(cfg)
	lst := NewListener(memBus, natsBus)
	// When
	ok := sdk.ImplementsIListener(lst)
	// Then
	assert.True(t, ok)

}
