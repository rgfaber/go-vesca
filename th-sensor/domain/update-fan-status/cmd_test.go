package update_fan_status

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	fan_status "github.com/rgfaber/go-vesca/th-sensor/model/fan-status"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	// Given
	id := model.NewTHSensorTestID()
	c := NewCmd(id.Id(), fan_status.Deactivated)
	// When
	// Then
	assert.NotNil(t, c)
}

func TestIfCmdImplementsICmd(t *testing.T) {
	// Given
	id := model.NewTHSensorTestID()
	c := NewCmd(id.Id(), fan_status.Deactivated)
	// When
	v, ok := interface{}(c).(dec.ICmd)
	// Then
	assert.NotNil(t, v)
	assert.True(t, ok)
}
