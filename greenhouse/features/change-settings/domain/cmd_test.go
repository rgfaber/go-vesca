package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	// Given
	id := model.BogusGreenhouseID
	settings := model.BogusSettings
	// And
	cmd := NewCmd(id, settings)
	// Then
	assert.NotNil(t, cmd)
	assert.NotNil(t, cmd.aggregateId)
	assert.NotNil(t, cmd.Settings)
}

func TestCmd_ImplementsICmd(t *testing.T) {
	// Given
	cmd := NewCmd(nil, nil)
	// When
	ok := sdk.ImplementsICmd(cmd)
	// Then
	assert.True(t, ok)

}
