package initialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {

	c := NewCmd(10.0, 10.0)
	assert.NotNil(t, c)
	assert.Equal(t, 10.0, c.measurement.Temperature)
	assert.Equal(t, 10.0, c.measurement.Humidity)
}
