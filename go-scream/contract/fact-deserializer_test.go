package contract

import (
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFactDeserializer(t *testing.T) {
	// Given

	// When
	fd := NewFactDeserializer(bogus.MyFactFromJson)
	// Then
	assert.NotNil(t, fd)
}

func TestFactDeserializer_ImplementsIFactDeserializer(t *testing.T) {
	// Given
	fd := 	NewFactDeserializer(nil)
	// When
	ok := interfaces.ImplementsIFactDeserializer(fd)
	// Then
	assert.True(t, ok)

}

func TestFactDeserializer_Deserialize(t *testing.T) {
	// Given
	mf := bogus.NewBogusFact("123", "abc")
	ser := NewFactSerializer(bogus.MyFactToJson)
	data, _ := ser.Serialize(mf)
	// And
	fd := NewFactDeserializer(bogus.MyFactFromJson)
	obj := fd.Deserialize(data)
	assert.NotNil(t, obj)
}
