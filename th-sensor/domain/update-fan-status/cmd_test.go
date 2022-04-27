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

func VerifyICmd(c dec.ICmd) bool {
	return true
}

func TestIfCmdImplementsICmd(t *testing.T) {
	// Given
	id := model.NewTHSensorTestID()
	c := NewCmd(id.Id(), fan_status.Deactivated)
	// When
	ok := VerifyICmd(c)
	// Then
	assert.True(t, ok)
}