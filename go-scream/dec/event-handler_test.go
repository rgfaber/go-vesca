package dec

import (
	mocks "github.com/rgfaber/go-vesca/sdk/infra/memory/mocks/interfaces"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewEventHandler(t *testing.T) {
	// Given
	membus := mocks.NewIDECBus(t)
	topic := "hello:there"
	// When
	eh := NewEventHandler(membus, "", nil)
	// Then
	assert.NotNil(t, eh)
	assert.Equal(t, membus, eh.MemBus)
	assert.Equal(t, topic, eh.Topic)
}

func Test_Implements_IEventHandler(t *testing.T) {
	// Given
	eh := NewEventHandler(nil, "", nil)
	// When
	ok := interfaces.ImplementsIEventHandler(eh)
	// Then
	assert.True(t, ok)
}
