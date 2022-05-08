package dec

import (
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFactHandler(t *testing.T) {
	// Given
	fs := NewFactDeserializer(bogus.MyFactFromJson)
	// When
	fh := NewFactHandler(fs, nil)
	// Then
	assert.NotNil(t, fh)
}

func TestFactHandler_Handle(t *testing.T) {
	// Given
	factHandled := false
	fn := func(fact interfaces.IFact) {
		factHandled = true
	}
	fd := NewFactDeserializer(bogus.MyFactFromJson)
	fh := NewFactHandler(fd, fn)
	fct := bogus.NewMyFact("aggregate", "trace")
	fs := NewFactSerializer(bogus.MyFactToJson)
	data, _ := fs.Serialize(fct)
	// When

	fh.Handle(data)
	// Then
	assert.True(t, factHandled)

}

func TestFactHandler_ImplementsIFactHandler(t *testing.T) {
	// Given
	fh := NewFactHandler(nil, nil)
	// When
	ok := interfaces.ImplementsIFactHandler(fh)
	// Then
	assert.True(t, ok)
}
