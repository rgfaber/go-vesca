package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMyHope(t *testing.T) {
	// Given
	// When
	mh := NewMyHope("aggregate", "trace")
	// Then
	assert.NotNil(t, mh)
	assert.Equal(t, "aggregate", mh.AggregateId())
	assert.Equal(t, "trace", mh.TraceId())
}

func TestMyHope_ImplementsIHope(t *testing.T) {
	// Given
	mh := NewMyHope("", "")
	// When
	ok := interfaces.ImplementsIHope(mh)
	// Then
	assert.True(t, ok)
}
