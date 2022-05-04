package nats

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListener(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	memBus := sdk.NewDECBus()
	natsBus := infra.NewNatsBus(cfg)
	// When
	lst := NewListener(memBus, natsBus)
	// Then
	assert.NotNil(t, lst)
}

func TestListenerImplementsIListener(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	memBus := sdk.NewDECBus()
	natsBus := infra.NewNatsBus(cfg)
	lst := NewListener(memBus, natsBus)
	// When
	ok := sdk.ImplementsIListener(lst)
	// Then
	assert.True(t, ok)


}


