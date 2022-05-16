package core

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/core/test"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewIdentity(t *testing.T) {
	// Given
	prefix := "prefix"
	// When
	id := NewIdentity(prefix)
	// Then
	assert.NotNil(t, id)

}

func TestNewDefaultIdentity(t *testing.T) {
	id := NewIdentity("")
	assert.NotNil(t, id)
	assert.Equal(t, "id", id.Prefix)
}

func TestNewIdentityFrom(t *testing.T) {
	// Given
	pf := test.TEST_PREFIX
	uuid := test.TEST_UUID
	// When
	id := NewIdentityFrom(pf, uuid)
	// Then
	assert.NotNil(t, id)
	assert.Equal(t, pf, id.Prefix)
	assert.Equal(t, test.CLEAN_TEST_UUID, id.Value)
	assert.Equal(t, fmt.Sprintf("%s-%s", test.TEST_PREFIX, test.CLEAN_TEST_UUID), id.Id())
}

func TestIdentity_Id(t *testing.T) {
	// Given
	testPrefix := test.TEST_PREFIX
	// When
	id := NewIdentity(testPrefix).Id()
	// Then
	assert.True(t, strings.HasPrefix(id, testPrefix))
}

func TestImplementsIIdentity(t *testing.T) {
	// Given
	testId := NewIdentity(test.TEST_PREFIX)
	// When
	b := interfaces.ImplementsIIdentity(testId)
	// Then
	assert.True(t, b)
}
