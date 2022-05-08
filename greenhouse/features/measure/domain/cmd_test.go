package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
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
	b := sdk.ImplementsICmd(cmd)
	// Then
	assert.True(t, b)
}
