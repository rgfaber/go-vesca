package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStore(t *testing.T) {
	// Given
	// When
	s := SingletonStore()
	// Then
	assert.NotNil(t, s)
}
