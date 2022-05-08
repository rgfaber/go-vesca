package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyFact_ImplementsIFact(t *testing.T) {
	// Given
	f := NewMyFact("", "")
	// When
	ok := interfaces.ImplementsIFact(f)
	// Then
	assert.True(t, ok)
}

func TestNewMyFact(t *testing.T) {
	// Given
	// When
	mf := NewMyFact("aggregate", "trace")
	// Then
	assert.NotNil(t, mf)
	assert.Equal(t, "aggregate", mf.AggregateId())
	assert.Equal(t, "trace", mf.TraceId())

}
