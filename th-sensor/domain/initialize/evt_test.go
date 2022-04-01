package initialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEvt(t *testing.T) {
	evt := NewEvt()
	assert.NotNil(t, evt)
}
