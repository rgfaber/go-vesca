package initialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChEvt(t *testing.T) {
	e := NewChEvt(10)
	assert.NotNil(t, e, "Could not create channel initialize.ChEvt")
}
