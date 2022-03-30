package initialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvents(t *testing.T) {
	e := NewEvents(10)
	assert.NotNil(t, e, "Could not create channel initialize.Events")
}
