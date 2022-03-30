package initialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCommands(t *testing.T) {
	c := NewCommands(100)
	assert.NotNil(t, c, "Could not create channel initialize.Commands")
}
