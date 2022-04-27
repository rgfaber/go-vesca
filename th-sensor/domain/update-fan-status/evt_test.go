package update_fan_status

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	fan_status "github.com/rgfaber/go-vesca/th-sensor/model/fan-status"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvt(t *testing.T) {
	// Given
	id := model.NewTHSensorTestID()
	e := NewEvt(id.Id(), fan_status.Activated)
	// When
	// Then
	assert.NotNil(t, e)
}

func TestIfEvtImplementsIEvt(t *testing.T) {
	// Given
	id := model.NewTHSensorTestID()
	e := NewEvt(id.Id(), fan_status.Activated)
	// When
	v, ok := interface{}(e).(dec.IEvt)
	// Then
	assert.NotNil(t, v)
	assert.True(t, ok)
}
