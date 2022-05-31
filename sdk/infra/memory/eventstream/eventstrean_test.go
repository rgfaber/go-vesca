package eventstream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewMemEventStream(t *testing.T) {
	// Given
	// When
	es := NewEventStream
	// Then
	assert.NotNil(t, es)
}
