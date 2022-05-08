package domain

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	// Given
	s := sdk.NewStore()
	b := sdk.NewDECBus()
	// When
	var a sdk.IAggregate = NewAggregate(s, b)
	// Then
	assert.NotNil(t, a)
}
