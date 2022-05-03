package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	bus := sdk.NewDECBus()
	store := infra.NewStore()
	a := NewAggregate(store, bus)
	assert.NotNil(t, a)
	assert.Nil(t, a.state)
}

func TestAggregate_AttemptInitializeCmdShouldResultInInitializedState(t *testing.T) {
	// Given
	bus := sdk.NewDECBus()
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
	bus := sdk.NewDECBus()
	store := infra.NewStore()
	a := NewAggregate(store, bus)
	// When
	b := sdk.ImplementsIAggregate(a)
	// Then
	assert.True(t, b)
}
