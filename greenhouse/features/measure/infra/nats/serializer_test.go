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
	s := NewSerializer()
	// Then
	assert.NotNil(t, s)
}

func TestSerializer_Serialize(t *testing.T) {
	// Given
	aggregateId := model.NewGreenhouseTestID().Id()
	traceId := sdk.TEST_TRACE_ID
	gh := model.NewTestGreenhouse()
	fact := contract.NewFact(aggregateId, traceId, gh)
	// AND
	ser := NewSerializer()
	// WHEN
	result := ser.Serialize(fact)
	// THEN
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)

}
