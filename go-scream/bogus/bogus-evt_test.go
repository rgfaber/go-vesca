package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewEvt(t *testing.T) {
	// Given
	id := core.NewIdentityFrom(BOGUS_PREFIX, testing2.TEST_UUID)
	sessionId := testing2.TEST_TRACE_ID
	name := "John Dow"
	age := 42
	weight := 82.3
	payload := NewBogusPayload(name, age, weight)
	// When
	evt := NewEvt(id, sessionId, payload)
	// Then
	assert.NotNil(t, evt)
	assert.Equal(t, id, evt.ID)
	assert.Equal(t, payload, evt.Payload)

}
