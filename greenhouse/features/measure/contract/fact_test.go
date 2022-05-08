package contract

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFact(t *testing.T) {
	// Given
	// When
	f := NewFact()
	// Then
	assert.NotNil(t, f)
}
