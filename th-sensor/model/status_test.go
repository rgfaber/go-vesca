package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSensorStatus_HasFlag(t *testing.T) {

}

func TestSensorStatus_Set(t *testing.T) {
	s := Unknown
	assert.True(t, s.HasFlag(Unknown))
	s = s.Set(Created)
	assert.True(t, s.HasFlag(Created))
	s = s.Set(Initialized)
	assert.True(t, s.HasFlag(Initialized))
	s = s.Set(Measuring)
	assert.True(t, s.HasFlag(Measuring))
	fmt.Println(s)
}

func TestSensorStatus_Unset(t *testing.T) {
	s := Created
	assert.True(t, s.HasFlag(Created))
	s = s.Unset(Created)
	assert.False(t, s.HasFlag(Created))
	assert.True(t, s.HasFlag(Unknown))
}
