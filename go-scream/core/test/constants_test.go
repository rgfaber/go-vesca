package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTEST_TRACE_ID(t *testing.T) {
	assert.Equal(t, "test_trace_id", TEST_TRACE_ID)
}
