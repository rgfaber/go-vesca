package initialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRsp(t *testing.T) {
	r := NewRsp("", "")
	assert.NotNil(t, r)
}
