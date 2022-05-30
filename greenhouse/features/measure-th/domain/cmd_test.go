package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	// Given
	id := model.NewGreenhouseID(testing2.TEST_UUID)
	traceId, _ := core.NewUuid()
	// When
	cmd := NewCmd(*id, traceId)
	// Then
	assert.NotNil(t, cmd)
}

func TestCmdImplementsICmd(t *testing.T) {
	// Given
	id := model.NewGreenhouseTestID()
	traceId := testing2.TEST_TRACE_ID
	cmd := NewCmd(*id, traceId)
	// When
	b := domain.ImplementsICmd(cmd)
	// Then
	assert.True(t, b)
}
