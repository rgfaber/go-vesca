package initialize

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRsp(t *testing.T) {
	id := NewTHSensorTestId
	r := NewRsp("", "")
	assert.NotNil(t, r)
}
