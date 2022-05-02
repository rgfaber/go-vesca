package measure

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	// Given
	id := model.NewGreenhouseID(sdk.TEST_UUID)
	traceId, _ := sdk.NewUuid()
	// When
	cmd := NewCmd(*id, traceId)
	// Then
	assert.NotNil(t, cmd)
}

func TestCmdImplementsICmd(t *testing.T) {
	// Given
	id := model.NewGreenhouseTestID()
	traceId := sdk.TEST_TRACE_ID
	cmd := NewCmd(*id, traceId)
	// When
	b := dec.ImplementsICmd(cmd)
	// Then
	assert.True(t, b)
}
