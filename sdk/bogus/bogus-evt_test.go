package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewEvt(t *testing.T) {
	// Given
	id := core.NewIdentityFrom(BOGUS_PREFIX, mocks.TEST_UUID)
	sessionId := mocks.TEST_TRACE_ID
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
