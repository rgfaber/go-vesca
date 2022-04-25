package sdk

import (
	"github.com/rgfaber/go-vesca/sdk/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMeta(t *testing.T) {
	id := NewIdentityFrom(constants.TEST_PREFIX, constants.TEST_UUID)
	m := NewMeta(*id, constants.TEST_TRACE_ID, 42)
	assert.NotNil(t, m)
	assert.NotNil(t, m.ID)
	assert.Equal(t, id.Id(), m.ID.Id())
}
