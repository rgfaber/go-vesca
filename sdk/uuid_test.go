package sdk

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNullId(t *testing.T) {
	// Given
	id := NullId()
	// When
	newId := strings.Replace(id, "0", "A", -1)
	assert.Equal(t, "AAAAAAAA-AAAA-AAAA-AAAA-AAAAAAAAAAAA", newId)
}
