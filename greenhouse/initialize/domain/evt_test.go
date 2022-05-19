package domain

import (
	testing2 "github.com/rgfaber/go-vesca/go-scream/core/test"
	"github.com/rgfaber/go-vesca/go-scream/interfaces"
	"github.com/rgfaber/go-vesca/greenhouse/initialize/test"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvt(t *testing.T) {
	// Given
	aggregateId := model.BogusGreenhouseID
	// When
	evt := NewEvt(aggregateId, testing2.TEST_TRACE_ID, nil)
	// Then
	assert.NotNil(t, evt)
}

func TestEvt_ImplementsIEvt(t *testing.T) {
	// Given
	e := test.BogusEvt
	// When
	b := interfaces.ImplementsIEvt(e)
	// Then
	assert.True(t, b)
}
