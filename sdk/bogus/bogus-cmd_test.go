package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewMyCmd(t *testing.T) {
	// Given
	id := core.NewIdentityFrom(mocks.TEST_PREFIX, mocks.TEST_UUID)
	// When
	cmd := NewBogusCmd(id, mocks.TEST_TRACE_ID, nil)
	// Then
	assert.NotNil(t, cmd)
}

func Test_MyCmd_ImplementsICmd(t *testing.T) {
	// Given
	cmd := NewBogusCmd(nil, mocks.TEST_TRACE_ID, nil)
	// When
	ok := domain.ImplementsICmd(cmd)
	// Then
	assert.True(t, ok)
}
