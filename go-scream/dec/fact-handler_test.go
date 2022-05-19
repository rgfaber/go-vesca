package dec

import (
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFactHandler(t *testing.T) {
	// Given
	// When
	fh := NewFactHandler(nil, nil)
	// Then
	assert.NotNil(t, fh)
}

func TestFactHandler_Handle(t *testing.T) {
	// Given
	factHandled := false
	handler := func(fact interfaces.IFact) {
		factHandled = true
	}
	fh := NewFactHandler(handler, bogus.MyFactFromJson)
	fct := bogus.NewBogusFact("aggregate", "trace")

	data, _ := bogus.MyFactToJson(fct)
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
