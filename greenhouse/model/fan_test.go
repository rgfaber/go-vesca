package model

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFan(t *testing.T) {
	// Given
	name := "fan 0"
	id := sdk.TEST_UUID
	// When
	fan := NewFan(id, name)
	// Then
	assert.NotNil(t, fan)
	assert.NotNil(t, fan.Details)
	assert.Equal(t, id, fan.ID.Value())
	assert.Equal(t, name, fan.Details.Name)
}
