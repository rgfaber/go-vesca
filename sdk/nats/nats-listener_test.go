package nats

import (
	initialize_infra_nats "github.com/rgfaber/go-vesca/greenhouse/features/initialize/infra/nats"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/rgfaber/go-vesca/sdk/dec"
	sdk_nats_config "github.com/rgfaber/go-vesca/sdk/nats/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewListener(t *testing.T) {
	// Given
	cfg := sdk_nats_config.NewNatsConfig()
	natsBus := NewNatsBus(cfg)
	startedSerializer = NewFactSerializer
	startedHandler = dec.NewFactHandler()
	// When
	lst := NewListener(natsBus, bogus.BOGUS_FACT_TOPIC, startedHandler)
	// Then
	assert.NotNil(t, lst)
}

func TestListenerImplementsIListener(t *testing.T) {
	// Given
	cfg := sdk_nats_config.NewNatsConfig()
	memBus := sdk.NewDECBus()
	natsBus := NewNatsBus(cfg)
	serializer := initialize_infra_nats.NewFactSerializer()
	lst := NewListener(memBus, natsBus, initialize_infra_nats.FACT_TOPIC, serializer)
	// When
	ok := sdk.ImplementsIListener(lst)
	// Then
	assert.True(t, ok)

}
