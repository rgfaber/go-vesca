package domain

import (
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewEventHandler(t *testing.T) {
	// Given
	membus := mediator.SingletonDECBus()
	topic := "hello:there"
	// When
	eh := NewEventHandler(membus, topic, nil)
	// Then
	assert.NotNil(t, eh)
	assert.Equal(t, membus, eh.MemBus)
	assert.Equal(t, topic, eh.Topic)
}

func Test_Implements_IEventHandler(t *testing.T) {
	// Given
	eh := NewEventHandler(nil, "", nil)
	// When
	ok := ImplementsIEventHandler(eh)
	// Then
	assert.True(t, ok)
}
