package eventstream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewMemEventStream(t *testing.T) {
	// Given
	// When
	es := NewMemEvents
	// Then
	assert.NotNil(t, es)
}
