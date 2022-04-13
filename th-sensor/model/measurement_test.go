package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMeasurement(t *testing.T) {
	m := NewMeasurement(15.0, 7.0)
	assert.NotNil(t, m)
	assert.Equal(t, 15.0, m.Temperature)
	assert.Equal(t, 7.0, m.Humidity)
}
