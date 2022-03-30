package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSensorStatus_HasFlag(t *testing.T) {

}

func TestSensorStatus_Set(t *testing.T) {
	s := SensorStatus(Unknown)
	assert.True(t, s.HasFlag(Unknown))
	s = s.Set(Created)
	assert.True(t, s.HasFlag(Created))
	s = s.Set(Initialized)
	assert.True(t, s.HasFlag(Initialized))
}

func TestSensorStatus_Unset(t *testing.T) {
	s := SensorStatus(Created)
	assert.True(t, s.HasFlag(Created))
	s = s.Unset(Created)
	assert.False(t, s.HasFlag(Created))
	assert.True(t, s.HasFlag(Unknown))
}