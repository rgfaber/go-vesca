package measure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewChEvt(t *testing.T) {
	ch := NewChEvt(10)
	assert.NotNil(t, ch)
	assert.Equal(t, cap(ch), 10)
}
