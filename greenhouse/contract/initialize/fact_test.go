package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFact(t *testing.T) {
	// Given
	aggregateId := model.NewGreenhouseTestID().Id()
	traceId := sdk.TEST_TRACE_ID
	gh :=
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
	traceId := sdk.TEST_TRACE_ID
	gh :=
	f := NewFact(aggregateId, traceId)
	// When
	ok := dec.ImplementsIFact(f)
	// Then
	assert.True(t, ok)
}
