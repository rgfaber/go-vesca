package update_fan_status

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	// Given
	store := infra.NewStore()
	bus := dec.NewDECBus()
	a := NewAggregate(store, bus)
	// When
	assert.NotNil(t, a)
}

func VerifyIAggregate(a dec.IAggregate) bool {
	return true
}

func TestIfAggregateImplementsIAggregate(t *testing.T) {
	// Given
	store := infra.NewStore()
	bus := dec.NewDECBus()
	a := NewAggregate(store, bus)
	// When
	ok := VerifyIAggregate(a)
	// Then
	assert.True(t, ok)
}
