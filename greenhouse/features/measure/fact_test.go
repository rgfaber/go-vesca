package measure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFact(t *testing.T) {
	// Given
	// When
	f := NewFact("", nil)
	// Then
	assert.NotNil(t, f)
}
