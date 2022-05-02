package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSettings(t *testing.T) {
	// Given
	temp := 15.0
	hum := 42.0
	// When
	m := NewSettings(temp, hum)
	// Then
	assert.NotNil(t, m)
	assert.Equal(t, temp, m.Temperature)
	assert.Equal(t, hum, m.Humidity)
}
