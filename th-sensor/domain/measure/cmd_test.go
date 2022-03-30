package measure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCmd(t *testing.T) {
	cmd := NewCmd()
	assert.NotNil(t, cmd)
}
