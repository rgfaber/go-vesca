package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGreenhouse(t *testing.T) {
	// Given
	// When
	gh := NewTestGreenhouse()
	// Then
	assert.NotNil(t, gh)
}
