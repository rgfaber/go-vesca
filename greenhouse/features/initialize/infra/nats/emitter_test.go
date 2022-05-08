package nats

import (
	"fmt"
	initialize_contract "github.com/rgfaber/go-vesca/greenhouse/features/initialize/contract"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/nats"
	sdk_nats_config "github.com/rgfaber/go-vesca/sdk/nats/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmitter(t *testing.T) {
	// Given
	cfg := sdk_nats_config.NewNatsConfig()
	nats := nats.NewNatsBus(cfg)
	bus := sdk.NewDECBus()
	ser := NewFactSerializer()
	// When
	emitter := NewEmitter(bus, nats, FACT_TOPIC, ser)
	// Then
	assert.NotNil(t, emitter)
}

func TestIfEmitterImplementsIEmitter(t *testing.T) {
	// Given
	cfg := sdk_nats_config.NewNatsConfig()
	bus := nats.NewNatsBus(cfg)
	memBus := sdk.NewDECBus()
	ser := NewFactSerializer()
	e := NewEmitter(memBus, bus, FACT_TOPIC, ser)
	// When
	ok := sdk.ImplementsIEmitter(e)
	//
	assert.True(t, ok)
}

func TestEmitter_Emit(t *testing.T) {
	// Given
	cfg := sdk_nats_config.NewNatsConfig()
	natsBus := nats.NewNatsBus(cfg)
	memBus := sdk.NewDECBus()
	aggregateId := model.NewGreenhouseTestID().Id()
	traceId := sdk.TEST_TRACE_ID
	gh := model.NewTestGreenhouse()
	ser := NewFactSerializer()
	e := NewEmitter(memBus, natsBus, FACT_TOPIC, ser)
	f := initialize_contract.NewFact(aggregateId, traceId, gh)
	receivedFact := false
	// AND
	ls := nats.NewListener(memBus, natsBus, FACT_TOPIC, ser)
	ls.Listen(func(fact sdk.IFact) {
		fc := fact.(*initialize_contract.Fact)
		receivedFact = true
		fmt.Printf("Received Fact [%v]\n on [%s]", fc, FACT_TOPIC)
	})
	// When
	e.Emit(f)
	// Then
	assert.True(t, receivedFact)

}
