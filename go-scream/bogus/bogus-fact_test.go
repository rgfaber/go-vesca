package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BogusFact_ImplementsIFact(t *testing.T) {
	// Given
	f := NewBogusFact("", "")
	// When
	ok := interfaces.ImplementsIFact(f)
	// Then
	assert.True(t, ok)
}

func Test_New_BogusFact(t *testing.T) {
	// Given
	// When
	mf := NewBogusFact("aggregate", "trace")
	// Then
	assert.NotNil(t, mf)
	assert.Equal(t, "aggregate", mf.AggregateId())
	assert.Equal(t, "trace", mf.TraceId())

}
