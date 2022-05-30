package domain

import (
	bogus2 "github.com/rgfaber/go-vesca/greenhouse/model/bogus"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvt(t *testing.T) {
	// Given
	aggregateId := bogus2.BogusGreenhouseID
	// When
	evt := NewEvt(aggregateId, testing2.TEST_TRACE_ID, nil)
	// Then
	assert.NotNil(t, evt)
}

func TestEvt_ImplementsIEvt(t *testing.T) {
	// Given
	e := BogusEvt
	// When
	b := domain.ImplementsIEvt(e)
	// Then
	assert.True(t, b)
}
