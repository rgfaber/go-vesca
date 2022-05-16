package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/helpers/infra"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	TestCfg = config.NewConfig()
	TestId  = model.NewGreenhouseTestID()
)

func TestNewAggregate(t *testing.T) {
	bus := mediator.SingletonDECBus()
	store := store.NewStore()
	a := NewAggregate(store, bus)
	assert.NotNil(t, a)
	assert.Nil(t, a.state)
}

func TestAggregate_AttemptInitializeCmdShouldResultInInitializedState(t *testing.T) {
	// Given
	bus := mediator.SingletonDECBus()
	store := store.NewStore()
	a := NewAggregate(store, bus)
	id := model.BogusGreenhouseID
	cmd := BogusCmd
	// And
	infra.SaveGreenhouse(store, model.BogusGreenhouse)
	// When
	fbk := a.Attempt(cmd)
	// Then
	assert.NotNil(t, fbk)
	assert.True(t, fbk.IsSuccess())
	m := store.Load(id.Id()).(*model.Greenhouse)
	assert.NotNil(t, m)
	assert.True(t, m.Status.HasFlag(model.Initialized))
}

func TestAggregateImplementsIAggregate(t *testing.T) {
	//Given
	bus := mediator.SingletonDECBus()
	store := store.NewStore()
	a := NewAggregate(store, bus)
	// When
	b := interfaces.ImplementsIAggregate(a)
	// Then
	assert.True(t, b)
}
