package measure

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func verifyIFbk(f dec.IFbk) bool {
	return true
}

func TestNewFbk(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	sensorId := cfg.SensorId()
	traceId := sdk.TEST_TRACE_ID
	isSuccess := false
	status := model.Initialized
	// When
	f := NewFbk(sensorId, traceId, isSuccess, status)
	// Then
	assert.NotNil(t, f)
	assert.Equal(t, sensorId, f.sensorId)
	assert.Equal(t, traceId, f.traceId)
	assert.Equal(t, isSuccess, f.isSuccess)
	assert.Equal(t, status, f.status)
}

func newTestFbk() *Fbk {
	cfg := config.NewConfig()
	sensorId := cfg.SensorId()
	traceId := sdk.TEST_TRACE_ID
	isSuccess := false
	status := model.Initialized
	return NewFbk(sensorId, traceId, isSuccess, status)
}

func TestIfFbkImplementsIFbk(t *testing.T) {
	// Given
	f := newTestFbk()
	// When
	b := verifyIFbk(f)
	// Then
	assert.True(t, b)
}
