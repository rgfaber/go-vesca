package initialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFact(t *testing.T) {
	f := NewFact()
	assert.NotNil(t, f)
}
