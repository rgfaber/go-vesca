package model

import (
	"github.com/rgfaber/go-vesca/greenhouse/model/bogus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGreenhouse(t *testing.T) {
	// Given
	// When
	gh := bogus.NewTestGreenhouse()
	// Then
	assert.NotNil(t, gh)
}
