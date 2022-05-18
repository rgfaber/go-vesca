package contract

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/test"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFact(t *testing.T) {
	// Given
	aggregateId := model.BogusGreenhouseID.Id()
	traceId := testing2.TEST_TRACE_ID
	gh := model.BogusGreenhouse

	// When
	f := NewFact(aggregateId, traceId, gh)
	// Then
	assert.NotNil(t, f)
	assert.Equal(t, aggregateId, f.AggregateId())
	assert.Equal(t, traceId, f.TraceId())
	assert.NotNil(t, f.Payload)
}

func TestIfFactImplementsIFact(t *testing.T) {
	// Given
	aggregateId := model.NewGreenhouseTestID().Id()
	traceId := testing2.TEST_TRACE_ID
	gh := model.NewTestGreenhouse()
	f := NewFact(aggregateId, traceId, gh)
	// When
	ok := interfaces.ImplementsIFact(f)
	// Then
	assert.True(t, ok)
}
