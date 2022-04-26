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

func TestCmdImplementsICmd(t *testing.T) {
	var c dec.ICmd = NewCmd()
}
