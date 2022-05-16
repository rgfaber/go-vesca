package infra

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveGreenhouse(t *testing.T) {
	// Given
	store := store.NewStore()
	state := model.NewTestGreenhouse()
	// And
	id := state.ID
	// When
	SaveGreenhouse(store, state)
	// And
	s := LoadGreenhouse(store, id.Id())
	// Then
	assert.NotNil(t, s)
}
