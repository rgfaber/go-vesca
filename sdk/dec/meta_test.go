package dec

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMeta(t *testing.T) {
	id := core.NewIdentityFrom(testing2.TEST_PREFIX, testing2.TEST_UUID)
	m := NewMeta(*id, testing2.TEST_TRACE_ID, 42)
	assert.NotNil(t, m)
	assert.NotNil(t, m.ID)
	assert.Equal(t, id.Id(), m.ID.Id())
}
