package initialize

import (
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvt(t *testing.T) {
	m := model.NewMeasurement(15.2, 50)
	aggregateId := model.NewTHSensorTestID()
	evt := NewEvt(aggregateId, TEST_TRACE_ID)
	assert.NotNil(t, evt)
}
