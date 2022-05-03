package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFact(t *testing.T) {
	// Given
	aggregateId := model.NewGreenhouseTestID().Id()
	traceId := sdk.TEST_TRACE_ID
	gh := model.NewTestGreenhouse()

	// When
	f := NewFact(aggregateId, traceId, gh)
	// Then
	assert.NotNil(t, f)
	assert.Equal(t, aggregateId, f.aggregateId)
	assert.Equal(t, traceId, f.traceId)
	assert.NotNil(t, f.payload)
}

func TestIfFactImplementsIFact(t *testing.T) {
	// Given
	aggregateId := model.NewGreenhouseTestID().Id()
	traceId := sdk.TEST_TRACE_ID
	gh := model.NewTestGreenhouse()
	f := NewFact(aggregateId, traceId, gh)
	// When
	ok := sdk.ImplementsIFact(f)
	// Then
	assert.True(t, ok)
}

func TestFact_Serialize(t *testing.T) {
	// Given
	aggregateId := model.NewGreenhouseTestID().Id()
	traceId := sdk.TEST_TRACE_ID
	gh := model.NewTestGreenhouse()
	f := NewFact(aggregateId, traceId, gh)
	// When
	s := f.Serialize(f)
	// Then
	assert.Equal(t, "", s)

}
