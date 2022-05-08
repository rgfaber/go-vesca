package sdk

import (
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStore(t *testing.T) {
	// Given
	// When
	s := NewStore()
	// Then
	assert.NotNil(t, s)
}

func TestIfWeCanSaveAndLoadTheModel(t *testing.T) {

	// Given
	s := NewStore()
	m := bogus.NewBogusModel(TEST_UUID)
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
