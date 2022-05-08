package dec

import (
	"github.com/rgfaber/go-vesca/sdk/bogus"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHopeHandler(t *testing.T) {
	// Given
	handled := false
	fn := func(hope interfaces.IHope) (interfaces.IFbk, error) {
		handled = true
		return nil, nil
	}
	hdes := NewHopeDeserializer(bogus.MyHopeFromJson)
	fser := NewFbkSerializer(bogus.MyFbkToJson)
	// And
	handler := NewHopeHandler(fn, hdes, fser)
	// When
	handler.Handle(nil)
	// Then
	assert.True(t, handled)
}
