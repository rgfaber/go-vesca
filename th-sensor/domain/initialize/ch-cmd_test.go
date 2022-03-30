package initialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChCmd(t *testing.T) {
	c := NewChCmd(100)
	assert.NotNil(t, c, "Could not create channel initialize.ChCmd")
}
