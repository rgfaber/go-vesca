package measure

import (
	testing2 "github.com/rgfaber/go-vesca/go-scream/core/test"
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFbk(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	sensorId := cfg.SensorId()
	traceId := testing2.TEST_TRACE_ID
	isSuccess := false
	status := model.Initialized
	// When
	f := NewFbk(sensorId, traceId, isSuccess, status)
	// Then
	assert.NotNil(t, f)
	assert.Equal(t, sensorId, f.Aggregate_Id)
	assert.Equal(t, traceId, f.Trace_Id)
	assert.Equal(t, isSuccess, f.isSuccess)
	assert.Equal(t, status, f.Greenhouse_Status)
}

func newTestFbk() *Fbk {
	cfg := config.NewConfig()
	sensorId := cfg.SensorId()
	traceId := testing2.TEST_TRACE_ID
	isSuccess := false
	status := model.Initialized
	return NewFbk(sensorId, traceId, isSuccess, status)
}

func TestIfFbkImplementsIFbk(t *testing.T) {
	// Given
	f := newTestFbk()
	// When
	b := sdk.ImplementsIFbk(f)
	// Then
	assert.True(t, b)
}
