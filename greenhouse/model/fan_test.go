package model

import (
	testing2 "github.com/rgfaber/go-vesca/go-scream/core/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFan(t *testing.T) {
	// Given
	name := "fan 0"
	id := testing2.TEST_UUID
	// When
	fan := NewFan(id, name)
	// Then
	assert.NotNil(t, fan)
	assert.NotNil(t, fan.Details)
	assert.Equal(t, testing2.CLEAN_TEST_UUID, fan.ID.Value)
	assert.Equal(t, name, fan.Details.Name)
}
