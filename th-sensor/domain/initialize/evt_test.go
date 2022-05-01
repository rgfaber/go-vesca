package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvt(t *testing.T) {
	m := model.NewMeasurement(15.2, 50)
	aggregateId := model.NewTHSensorTestID()
	evt := NewEvt(aggregateId, sdk.TEST_TRACE_ID, *m)
	assert.NotNil(t, evt)
}

func TestEvt_ImplementsIEvt(t *testing.T) {
	// Given
	aggregateId := model.NewTHSensorTestID()
	e := NewEvt(aggregateId, sdk.TEST_TRACE_ID, *model.NewMeasurement(15.0, 42.0))
	// When
	b := dec.ImplementsIEvt(e)
	// Then
	assert.True(t, b)

}
