package sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMeta(t *testing.T) {
	id := NewIdentityFrom(TEST_PREFIX, TEST_UUID)
	m := NewMeta(*id, TEST_TRACE_ID, 42)
	assert.NotNil(t, m)
	assert.NotNil(t, m.ID)
	assert.Equal(t, id.Id(), m.ID.Id())
}
