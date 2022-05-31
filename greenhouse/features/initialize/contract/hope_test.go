package contract

import (
	bogus "github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/contract"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHope(t *testing.T) {
	// Given
	gh := bogus.BogusGreenhouse
	// When
	h := NewHope(gh.ID.Value, testing2.TEST_TRACE_ID, nil, nil, nil, nil, nil)
	// Then
	assert.NotNil(t, h)
}

func TestHopeImplementsIHope(t *testing.T) {
	// Given
	gh := bogus.BogusGreenhouse
	h := NewHope(gh.ID.Value, testing2.TEST_TRACE_ID, nil, nil, nil, nil, nil)
	// When
	ok := contract.ImplementsIHope(h)
	// Then
	assert.True(t, ok)
}
