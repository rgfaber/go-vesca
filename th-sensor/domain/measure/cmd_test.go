package measure

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	cmd := NewCmd()
	assert.NotNil(t, cmd)
}

func verifyICmd(c dec.ICmd) bool {
	return true
}

func TestCmdImplementsICmd(t *testing.T) {
	// Given
	cmd := NewCmd()
	// When
	b := verifyICmd(cmd)
	// Then
	assert.True(t, b)
}
