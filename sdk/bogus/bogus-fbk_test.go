package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFbk(t *testing.T) {
	// Given
	// When
	fbk := NewBogusFbk("aggregate", "trace", "ErrorMsg")
	// Then
	assert.NotNil(t, fbk)
	assert.Equal(t, "aggregate", fbk.OriginId)
	assert.Equal(t, "trace", fbk.SessionId)
	assert.Equal(t, "ErrorMsg", fbk.ErrorMsg)
}

func TestMyFbk_ImplementsIFbk(t *testing.T) {
	// Given
	fbk := NewBogusFbk("", "", "")
	// when
	ok := domain.ImplementsIFbk(fbk)
	// Then
	assert.True(t, ok)
}
