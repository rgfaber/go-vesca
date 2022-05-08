package nats

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/contract"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/nats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmitter(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	nats := nats.NewNatsBus(cfg)
	bus := sdk.NewDECBus()
	// When
	emitter := NewEmitter(bus, nats)
	// Then
	assert.NotNil(t, emitter)
}

func TestIfEmitterImplementsIEmitter(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	bus := nats.NewNatsBus(cfg)
	memBus := sdk.NewDECBus()
	e := NewEmitter(memBus, bus)
	// When
	ok := sdk.ImplementsIEmitter(e)
	//
	assert.True(t, ok)
}

func TestEmitter_Emit(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	natsBus := nats.NewNatsBus(cfg)
	memBus := sdk.NewDECBus()
	aggregateId := model.NewGreenhouseTestID().Id()
	traceId := sdk.TEST_TRACE_ID
	gh := model.NewTestGreenhouse()
	e := NewEmitter(memBus, natsBus)
	f := contract.NewFact(aggregateId, traceId, gh)
	// AND
	natsBus.Listen(e.Topic(), func(fact contract.Fact) {

	})

}
