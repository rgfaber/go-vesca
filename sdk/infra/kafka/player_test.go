package kafka

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewPlayer(t *testing.T) {
	// Given
	p := NewPlayer()
	// When
	// Then
	assert.NotNil(t, p)
}
