package dec

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New_Feature(t *testing.T) {
	// Given
	responders := interfaces.BuildResponders()
	listeners := interfaces.BuildListeners()
	handlers := interfaces.BuildHandlers()
	// When
	feature := NewFeature(responders, listeners, handlers)
	// Then
	assert.NotNil(t, feature)
	assert.NotNil(t, feature.Responders)
	assert.NotNil(t, feature.Listeners)
	assert.NotNil(t, feature.Handlers)
}

func Test_Feature_ImplementsIFeature(t *testing.T) {
	// Given
	responders := interfaces.BuildResponders()
	listeners := interfaces.BuildListeners()
	handlers := interfaces.BuildHandlers()
	feature := NewFeature(responders, listeners, handlers)
	// When
	ok := interfaces.ImplementsIFeature(feature)
	// Then
	assert.True(t, ok)
}
