package initialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	c := NewCmd()
	assert.NotNil(t, c, "Could not create initialize.Cmd")
}
