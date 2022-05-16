package domain

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	// Given
	store := store.NewStore()
	memBus := mediator.SingletonDECBus()
	// When
	agg := NewAggregate(memBus, store)
	// Then
	assert.NotNil(t, agg)
}

func TestAggregate_ImplementsIAggregate(t *testing.T) {
	// Given
	store := store.NewStore()
	memBus := mediator.SingletonDECBus()
	agg := NewAggregate(memBus, store)
	// When
	ok := sdk.ImplementsIAggregate(agg)
	// Then
	assert.True(t, ok)

}
