package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvt(t *testing.T) {
	// Given
	m := model.BogusSettings
	aggregateId := model.BogusGreenhouseID
	// When
	evt := NewEvt(aggregateId, sdk.TEST_TRACE_ID, m)
	// Then
	assert.NotNil(t, evt)
}

func TestEvt_ImplementsIEvt(t *testing.T) {
	// Given
	e := BogusEvt
	// When
	b := sdk.ImplementsIEvt(e)
	// Then
	assert.True(t, b)
}
