package integration

import (
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfWeCanSaveAndLoadTheModel(t *testing.T) {
	// Given
	s := store.SingletonStore()
	m := bogus.NewBogusModel(mocks.TEST_UUID)
	bogus.SaveBogusState(s, m)
	m.Status = bogus.Started
	m = bogus.LoadBogusState(s, m.ID.Id())
	assert.NotNil(t, m)
	assert.Equal(t, bogus.Created, m.Status)
	m.Status = bogus.Started
	bogus.SaveBogusState(s, m)
	m = bogus.LoadBogusState(s, m.ID.Id())
	assert.NotNil(t, m)
	assert.Equal(t, bogus.Started, m.Status)
}
