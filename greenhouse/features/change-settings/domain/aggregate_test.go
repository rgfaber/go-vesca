package domain

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	// Given
	store := sdk.NewStore()
	memBus := sdk.NewDECBus()
	// When
	agg := NewAggregate(memBus, store)
	// Then
	assert.NotNil(t, agg)
}

func TestAggregate_ImplementsIAggregate(t *testing.T) {
	// Given
	store := sdk.NewStore()
	memBus := sdk.NewDECBus()
	agg := NewAggregate(memBus, store)
	// When
	ok := sdk.ImplementsIAggregate(agg)
	// Then
	assert.True(t, ok)

}
