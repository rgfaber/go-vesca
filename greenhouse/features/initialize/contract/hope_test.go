package contract

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHope(t *testing.T) {
	// Given
	gh := model.BogusGreenhouse
	// When
	h := NewHope(gh.ID.Value, sdk.TEST_TRACE_ID, gh.Details.Name)
	// Then
	assert.NotNil(t, h)
	assert.Equal(t, gh.ID.Value, h.aggregateId)
	assert.Equal(t, gh.Details.Name, h.Name)
}

func TestHopeImplementsIHope(t *testing.T) {
	// Given
	h := NewHope("", "", "")
	// When
	ok := sdk.ImplementsIHope(h)
	// Then
	assert.True(t, ok)
}
