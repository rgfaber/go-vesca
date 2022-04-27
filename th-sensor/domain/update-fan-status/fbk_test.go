package update_fan_status

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	fan_status "github.com/rgfaber/go-vesca/th-sensor/model/fan-status"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFbk(t *testing.T) {
	// Given
	id := model.NewTHSensorTestID()
	f := NewFbk(id.Id(), sdk.TEST_TRACE_ID, true, fan_status.Activated)
	// When
	// Then
	assert.NotNil(t, f)
	assert.Equal(t, id.Id(), f.sensorId)
	assert.Equal(t, sdk.TEST_TRACE_ID, f.traceId)
	assert.Equal(t, true, f.isSuccess)
	assert.Equal(t, fan_status.Activated, f.status)
}
