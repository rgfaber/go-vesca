package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFact(t *testing.T) {
	// Given
	aggregateId := model.NewTHSensorTestID().Id()
	traceId := sdk.TEST_TRACE_ID
	// When
	f := NewFact(aggregateId, traceId)
	// Then
	assert.NotNil(t, f)
	assert.Equal(t, aggregateId, f.aggregateId)
	assert.Equal(t, traceId, f.traceId)
	assert.Nil(t, f.payload)
}

func TestIfFactImplementsIFact(t *testing.T) {
	// Given
	aggregateId := model.NewTHSensorTestID().Id()
	traceId := sdk.TEST_TRACE_ID
	f := NewFact(aggregateId, traceId)
	// When
	ok := dec.ImplementsIFact(f)
	// Then
	assert.True(t, ok)

}
