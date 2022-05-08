package nats

import (
	"github.com/rgfaber/go-vesca/greenhouse/features/initialize/contract"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSerializer(t *testing.T) {
	// Given
	// When
	s := NewFactSerializer()
	// Then
	assert.NotNil(t, s)
}

func TestSerializerImplementsISerializer(t *testing.T) {
	// Given
	s := NewFactSerializer()
	// When
	ok := sdk.ImplementsIFactSerializer(s)
	// Then
	assert.True(t, ok)

}

func TestSerializer_Serialize(t *testing.T) {
	// Given
	aggregateId := model.NewGreenhouseTestID().Id()
	traceId := sdk.TEST_TRACE_ID
	gh := model.NewTestGreenhouse()
	fact := contract.NewFact(aggregateId, traceId, gh)
	// AND
	ser := NewFactSerializer()
	// WHEN
	result, err := ser.Serialize(fact)
	// THEN
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)

}
