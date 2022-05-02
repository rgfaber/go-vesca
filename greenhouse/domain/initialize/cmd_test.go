package initialize

import (
	"fmt"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	id := model.NewGreenhouseTestID()
	c := NewCmd(*id, "Test Sensor", "00000", "test_trace_id", 10.0, 10.0)
	assert.NotNil(t, c)
	assert.Equal(t, 10.0, c.settings.Temperature)
	assert.Equal(t, 10.0, c.settings.Humidity)
}

func TestCmd_ID(t *testing.T) {
	id := model.NewGreenhouseTestID()
	c := NewCmd(*id, "Test Sensor", "00000", "test_trace_id", 10.0, 10.0)
	fmt.Println(c.AggregateId().Id())
}
