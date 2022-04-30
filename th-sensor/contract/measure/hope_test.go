package measure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHope(t *testing.T) {
	h := NewHope()
	assert.NotNil(t, h)
}
