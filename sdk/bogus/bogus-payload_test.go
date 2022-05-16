package bogus

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New_BogusPayload(t *testing.T) {
	// Given
	name := "John Dow"
	age := 25
	weight := 78.5
	// When
	mp := NewBogusPayload(name, age, weight)
	// Then
	assert.NotNil(t, mp)
	assert.Equal(t, name, mp.Name)
	assert.Equal(t, age, mp.Age)
	assert.Equal(t, weight, mp.Weight)
}
