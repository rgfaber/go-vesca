package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	id := model.NewTHSensorTestID()
	c := NewCmd(id.Value, "Test Sensor", "00000", "test_trace_id", 10.0, 10.0)
	assert.NotNil(t, c)
	assert.Equal(t, 10.0, c.measurement.Temperature)
	assert.Equal(t, 10.0, c.measurement.Humidity)
}

func TestCmd_ID(t *testing.T) {
	id := model.NewTHSensorTestID()
	c := NewCmd(id.Value, "Test Sensor", "00000", "test_trace_id", 10.0, 10.0)
	aggregateId := c.ID()
	fmt.Println(aggregateId)
}
