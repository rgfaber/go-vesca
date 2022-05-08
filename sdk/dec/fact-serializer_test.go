package dec

import (
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFactSerializer(t *testing.T) {
	// Given
	fact := bogus.NewMyFact("aggregate", "trace")
	fs := NewFactSerializer(bogus.MyFactToJson)
	// When
	s, err := fs.Serialize(fact)
	// Then
	assert.Nil(t, err)
	assert.NotEmpty(t, s)
}

func Test_FactSerializer_ImplementsIFactSerializer(t *testing.T) {
	// Given
	fs := NewFactSerializer(nil)
	// When
	ok := interfaces.ImplementsIFactSerializer(fs)
	// Then
	assert.True(t, ok)
}
