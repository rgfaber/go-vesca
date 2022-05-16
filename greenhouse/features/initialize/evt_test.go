package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/test"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
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
	e := BogusEvt
	// When
	b := interfaces.ImplementsIEvt(e)
	// Then
	assert.True(t, b)
}
