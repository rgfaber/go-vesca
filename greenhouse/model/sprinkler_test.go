package model

import (
	testing2 "github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSprinkler(t *testing.T) {
	// Given
	sprinklerName := "Sprinkler 0"
	id := testing2.TEST_UUID
	// When
	spr := NewSprinkler(id, sprinklerName)
	// Then
	assert.NotNil(t, spr)
	assert.Equal(t, testing2.CLEAN_TEST_UUID, spr.ID.Value)
	assert.Equal(t, sprinklerName, spr.Details.Name)
}
