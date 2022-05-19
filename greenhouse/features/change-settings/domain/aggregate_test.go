package domain

import (
	"github.com/rgfaber/go-vesca/go-scream/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/go-scream/infra/memory/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	// Given
	store := store.SingletonStore()
	memBus := mediator.SingletonDECBus()
	// When
	agg := NewAggregate(memBus, store)
	// Then
	assert.NotNil(t, agg)
}

func TestAggregate_ImplementsIAggregate(t *testing.T) {
	// Given
	store := store.SingletonStore()
	memBus := mediator.SingletonDECBus()
	agg := NewAggregate(memBus, store)
	// When
	ok := sdk.ImplementsIAggregate(agg)
	// Then
	assert.True(t, ok)

}
