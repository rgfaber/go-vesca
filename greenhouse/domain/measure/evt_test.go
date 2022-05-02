package measure

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvt(t *testing.T) {
	// Given
	id := model.NewGreenhouseTestID()
	sensorId := id.Value
	traceId := sdk.TEST_TRACE_ID
	temperature := 15.0
	humidity := 30.0

	// When
	evt := NewEvt(id, traceId, temperature, humidity)
	// Then
	assert.NotNil(t, evt)
	assert.Equal(t, sensorId, evt.aggregateId.Id())
	assert.Equal(t, temperature, evt.temperature)
	assert.Equal(t, humidity, evt.humidity)
}

func verifyIEvt(e dec.IEvt) bool {
	return true
}

func TestIfEvtImplementsIEvt(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	id := model.NewGreenhouseID(cfg.GreenhouseId())
	traceId := sdk.TEST_TRACE_ID
	temp := 30.0
	hum := 22.0
	evt := NewEvt(id, traceId, temp, hum)
	// When
	b := verifyIEvt(evt)
	// Then
	assert.True(t, b)
}
