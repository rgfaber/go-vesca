package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewTransientHttpResponder(t *testing.T) {
	// Given
	// When
	rsp := TransientHttpResponder()
	// Then
	assert.NotNil(t, rsp)

}
