package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	bus := dec.NewDECBus()
	store := infra.NewStore()
	a := NewAggregate(store, bus)
	assert.NotNil(t, a)
	assert.Nil(t, a.state)
}

func TestAggregate_AttemptInitializeCmdShouldResultInInitializedState(t *testing.T) {
	// Given
	bus := dec.NewDECBus()
	store := infra.NewStore()
	a := NewAggregate(store, bus)
	cfg := config.NewConfig()
	id := model.NewGreenhouseTestID()
	cmd := NewCmd(*id, cfg.SensorName(), cfg.GreenhouseId(), sdk.TEST_TRACE_ID, 15.0, 50.0)
	// When
	rsp, e := a.Attempt(cmd)
	// Then
	assert.Nil(t, e)
	assert.NotNil(t, rsp)
	m := store.Load(id.Id()).(*model.Greenhouse)
	assert.NotNil(t, m)
	assert.True(t, m.Status.HasFlag(model.Initialized))
}

func TestAggregateImplementsIAggregate(t *testing.T) {
	//Given
	bus := dec.NewDECBus()
	store := infra.NewStore()
	a := NewAggregate(store, bus)
	// When
	b := dec.ImplementsIAggregate(a)
	// Then
	assert.True(t, b)
}
