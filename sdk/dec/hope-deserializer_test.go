package dec

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHopeDeserializer(t *testing.T) {
	// Given
	// When
	hd := NewHopeDeserializer(nil)
	// Then
	assert.NotNil(t, hd)
}
