package store

import (
	"github.com/rgfaber/go-vesca/sdk/bogus"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStore(t *testing.T) {
	// Given
	// When
	s := SingletonStore()
	// Then
	assert.NotNil(t, s)
}

func TestIfWeCanSaveAndLoadTheModel(t *testing.T) {

	// Given
	s := SingletonStore()
	m := bogus.NewBogusModel(testing2.TEST_UUID)
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
