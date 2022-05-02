package initialize

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFbk(t *testing.T) {
	id := model.NewGreenhouseTestID()
	r := NewFbk(id.Id(), model.Initialized, true, "test")
	assert.NotNil(t, r)
}

func TestIfFbkImplementsIFbk(t *testing.T) {
	// Given
	id := model.NewGreenhouseTestID()
	fbk := NewFbk(id.Id(), model.Initialized, true, "test")
	// When
	c, ok := interface{}(fbk).(dec.IFbk)
	// Then
	assert.NotNil(t, c)
	assert.True(t, ok)

}