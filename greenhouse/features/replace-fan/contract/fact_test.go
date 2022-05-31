package contract

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	bogus "github.com/rgfaber/go-vesca/greenhouse/model"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFact(t *testing.T) {
	// Given
	aggregateId := bogus.BogusGreenhouseID.Id()
	traceId := testing2.TEST_TRACE_ID

	fan := bogus.BogusFan

	// When
	f := NewFact(aggregateId, traceId, fan)
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
	fan := bogus.BogusFan
	f := NewFact(aggregateId, traceId, fan)
	// When
	ok := domain.ImplementsIFact(f)
	// Then
	assert.True(t, ok)
}
