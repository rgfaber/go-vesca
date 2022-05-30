package core

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/core/mocks"
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
	pf := mocks.TEST_PREFIX
	uuid := mocks.TEST_UUID
	// When
	id := NewIdentityFrom(pf, uuid)
	// Then
	assert.NotNil(t, id)
	assert.Equal(t, pf, id.Prefix)
	assert.Equal(t, mocks.CLEAN_TEST_UUID, id.Value)
	assert.Equal(t, fmt.Sprintf("%s-%s", mocks.TEST_PREFIX, mocks.CLEAN_TEST_UUID), id.Id())
}

func TestIdentity_Id(t *testing.T) {
	// Given
	testPrefix := mocks.TEST_PREFIX
	// When
	id := NewIdentity(testPrefix).Id()
	// Then
	assert.True(t, strings.HasPrefix(id, testPrefix))
}

func TestImplementsIIdentity(t *testing.T) {
	// Given
	testId := NewIdentity(mocks.TEST_PREFIX)
	// When
	b := ImplementsIIdentity(testId)
	// Then
	assert.True(t, b)
}
