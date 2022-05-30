package core

import (
	core_mocks "github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMeta(t *testing.T) {
	id := NewIdentityFrom(core_mocks.TEST_PREFIX, core_mocks.TEST_UUID)
	m := NewMeta(id, core_mocks.TEST_TRACE_ID, 42)
	assert.NotNil(t, m)
	assert.NotNil(t, m.ID)
	assert.Equal(t, id.Id(), m.ID.Id())
}
