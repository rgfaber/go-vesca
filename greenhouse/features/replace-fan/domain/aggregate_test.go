package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/helpers/infra"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	bogus2 "github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	TestCfg = config.NewConfig()
	TestId  = model.NewGreenhouseTestID()
)

func TestNewAggregate(t *testing.T) {
	bus := mediator.SingletonDECBus()
	store := store.SingletonStore()
	a := NewAggregate(store, bus)
	assert.NotNil(t, a)
	assert.Nil(t, a.state)
}

func TestAggregate_AttemptInitializeCmdShouldResultInInitializedState(t *testing.T) {
	// Given
	bus := mediator.SingletonDECBus()
	store := store.SingletonStore()
	a := NewAggregate(store, bus)
	id := bogus2.BogusGreenhouseID
	cmd := BogusCmd
	// And
	infra.SaveGreenhouse(store, bogus2.BogusGreenhouse)
	// When
	fbk, _ := a.Attempt(cmd)
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
	store := store.SingletonStore()
	a := NewAggregate(store, bus)
	// When
	b := domain.ImplementsIAggregate(a)
	// Then
	assert.True(t, b)
}
