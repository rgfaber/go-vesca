package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/rgfaber/go-vesca/sdk/domain"
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
	f := NewFbk(sensorId, traceId, status, "")
	// Then
	assert.NotNil(t, f)
	assert.Equal(t, sensorId, f.Aggregate_Id)
	assert.Equal(t, traceId, f.Trace_Id)
	assert.Equal(t, isSuccess, f.IsSuccess())
	assert.Equal(t, status, f.Greenhouse_Status)
}

func newTestFbk() *Fbk {
	cfg := config.NewConfig()
	sensorId := cfg.SensorId()
	traceId := testing2.TEST_TRACE_ID
	status := model.Initialized
	return NewFbk(sensorId, traceId, status, "error")
}

func TestIfFbkImplementsIFbk(t *testing.T) {
	// Given
	f := newTestFbk()
	// When
	b := domain.ImplementsIFbk(f)
	// Then
	assert.True(t, b)
}
