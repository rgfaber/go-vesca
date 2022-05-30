package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFbk(t *testing.T) {
	id := model.NewGreenhouseTestID()
	r := NewFbk(id.Id(), model.Initialized, "", "test")
	assert.NotNil(t, r)
}

func TestIfFbkImplementsIFbk(t *testing.T) {
	// Given
	id := model.NewGreenhouseTestID()
	fbk := NewFbk(id.Id(), model.Initialized, "true", "test")
	// When
	ok := domain.ImplementsIFbk(fbk)
	// Then
	assert.True(t, ok)

}
