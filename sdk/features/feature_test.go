package features

import (
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/rgfaber/go-vesca/sdk/infra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New_Feature(t *testing.T) {
	// Given
	responders := infra.BuildResponders()
	listeners := infra.BuildListeners()
	handlers := domain.BuildHandlers()
	// When
	feature := NewFeature(responders, listeners, handlers)
	// Then
	assert.NotNil(t, feature)
	assert.Nil(t, feature.Responders)
	assert.Nil(t, feature.Listeners)
	assert.Nil(t, feature.Handlers)
}

func Test_Feature_ImplementsIFeature(t *testing.T) {
	// Given
	responders := infra.BuildResponders()
	listeners := infra.BuildListeners()
	handlers := domain.BuildHandlers()
	feature := NewFeature(responders, listeners, handlers)
	// When
	ok := ImplementsIFeature(feature)
	// Then
	assert.True(t, ok)
}
