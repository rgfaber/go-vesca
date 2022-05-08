package dec

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFbkSerializer(t *testing.T) {
	// Given
	fbkSer := NewFbkSerializer(nil)
	// When
	// Then
	assert.NotNil(t, fbkSer)
}

func Test_FbkSerializer_ImplementsIFbkSerializer(t *testing.T) {
	// Given
	fs := NewFbkSerializer(nil)
	// When
	ok := interfaces.ImplementsIFbkSerializer(fs)
	// Then
	assert.True(t, ok)
}
