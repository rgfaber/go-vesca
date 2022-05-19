package contract

import (
	testing2 "github.com/rgfaber/go-vesca/go-scream/core/test"
	"github.com/rgfaber/go-vesca/go-scream/interfaces"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHope(t *testing.T) {
	// Given
	gh := model.BogusGreenhouse
	// When
	h := NewHope(gh.ID.Value, testing2.TEST_TRACE_ID, nil, nil, nil, nil, nil)
	// Then
	assert.NotNil(t, h)
}

func TestHopeImplementsIHope(t *testing.T) {
	// Given
	gh := model.BogusGreenhouse
	h := NewHope(gh.ID.Value, testing2.TEST_TRACE_ID, nil, nil, nil, nil, nil)
	// When
	ok := interfaces.ImplementsIHope(h)
	// Then
	assert.True(t, ok)
}
