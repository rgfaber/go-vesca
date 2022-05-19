package measure

import (
	"github.com/rgfaber/go-vesca/go-scream/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/go-scream/infra/memory/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	// Given
	s := store.SingletonStore()
	b := mediator.SingletonDECBus()
	// When
	var a sdk.IAggregate = NewAggregate(s, b)
	// Then
	assert.NotNil(t, a)
}
