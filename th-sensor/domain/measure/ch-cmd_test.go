package measure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChCmd(t *testing.T) {
	chIn := NewChCmd(10)
	assert.NotNil(t, chIn)
	assert.Equal(t, cap(chIn), 10)
}
