package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSprinkler(t *testing.T) {
	// Given
	sprinklerName := "Sprinkler 0"
	id := sdk.TEST_UUID
	// When
	spr := NewSprinkler(id, sprinklerName)
	// Then
	assert.NotNil(t, spr)
	assert.Equal(t, sdk.CLEAN_TEST_UUID, spr.ID.Value)
	assert.Equal(t, sprinklerName, spr.Details.Name)
}
