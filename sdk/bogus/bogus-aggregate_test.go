package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New_BogusAggregate(t *testing.T) {
	// Given
	// When
	agg := NewBogusAggregate()
	// Then
	assert.NotNil(t, agg)
}

func TestMyAggregate_ImplementsIAggregate(t *testing.T) {
	// Given
	agg := NewBogusAggregate()
	// When
	ok := domain.ImplementsIAggregate(agg)
	// Then
	assert.True(t, ok)
}
