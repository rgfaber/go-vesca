package measure

import (
	"github.com/rgfaber/go-vesca/go-scream/core"
	testing2 "github.com/rgfaber/go-vesca/go-scream/core/test"
	"github.com/rgfaber/go-vesca/go-scream/interfaces"
	"github.com/rgfaber/go-vesca/greenhouse/model"
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
	b := interfaces.ImplementsICmd(cmd)
	// Then
	assert.True(t, b)
}
