package kill

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChCmd(t *testing.T) {
	ch := NewChCmd(10)
	assert.NotNil(t, ch)
	assert.Equal(t, cap(ch), 10)
}
