package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFbk(t *testing.T) {
	// Given
	// When
	fbk := NewMyFbk("aggregate", "trace", "error")
	// Then
	assert.NotNil(t, fbk)
	assert.Equal(t, "aggregate", fbk.aggregateId)
	assert.Equal(t, "trace", fbk.traceId)
	assert.Equal(t, "error", fbk.error)
}

func TestMyFbk_ImplementsIFbk(t *testing.T) {
	// Given
	fbk := NewMyFbk("", "", "")
	// when
	ok := interfaces.ImplementsIFbk(fbk)
	// Then
	assert.True(t, ok)
}
