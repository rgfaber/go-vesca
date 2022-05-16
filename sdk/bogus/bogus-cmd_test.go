package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/test"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewMyCmd(t *testing.T) {
	// Given
	id := core.NewIdentityFrom(testing2.TEST_PREFIX, testing2.TEST_UUID)
	// When
	cmd := NewBogusCmd(id, testing2.TEST_TRACE_ID, nil)
	// Then
	assert.NotNil(t, cmd)
}

func Test_MyCmd_ImplementsICmd(t *testing.T) {
	// Given
	cmd := NewBogusCmd(nil, testing2.TEST_TRACE_ID, nil)
	// When
	ok := interfaces.ImplementsICmd(cmd)
	// Then
	assert.True(t, ok)
}
