package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGreenhouse(t *testing.T) {
	// Given
	// When
	gh := NewTestGreenhouse()
	// Then
	assert.NotNil(t, gh)
}

func NewTestGreenhouse() *Greenhouse {
	id := NewGreenhouseTestID().Value()
	name := "Test Greenhouse"
	temp := 20.0
	hum := 42.0
	settings := NewSettings(temp, hum)
	sensorId := sdk.TEST_UUID
	sensorName := "Test Sensor"
	sensor := NewSensor(sensorId, sensorName)
	fanId := sdk.TEST_UUID
	fanName := "Test Fan"
	fan := NewFan(fanId, fanName)
	sprinklerId := sdk.TEST_UUID
	sprinklerName := "Test Sprinkler"
	sprinkler := NewSprinkler(sprinklerId, sprinklerName)
	return NewGreenhouse(id, name, settings, sensor, fan, sprinkler)
}
