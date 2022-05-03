package sdk

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewIdentity(t *testing.T) {
	id := NewIdentity(TEST_PREFIX)
	if id == nil {
		t.Errorf("id was not created")
	}
	if id.Prefix != TEST_PREFIX {
		t.Errorf("Identity has wrong Prefix. Expected 'my' got %+v", id.Prefix)
	}
}

func TestNewDefaultIdentity(t *testing.T) {
	id := NewIdentity("")
	assert.NotNil(t, id)
	assert.Equal(t, "id", id.Prefix)
}

func TestNewIdentityFrom(t *testing.T) {
	id := NewIdentityFrom(TEST_PREFIX, TEST_UUID)
	if id == nil {
		t.Errorf("No Identity was created")
	}
	if id.Prefix != TEST_PREFIX {
		t.Errorf("Identity has an incorrect Prefix. Expected [%+v] Got [%+v]", TEST_PREFIX, id.Prefix)
	}
	seed := strings.Replace(TEST_UUID, "-", "", -1)
	if id.Value != seed {
		t.Errorf("Identity has an incorrect Value. Expected [%+v] Got [%+v]", seed, id.Value)
	}
}

func TestIdentity_Id(t *testing.T) {
	// Given
	testPrefix := TEST_PREFIX
	// When
	id := NewIdentity(testPrefix).Id()
	// Then
	if !strings.HasPrefix(id, testPrefix) {
		t.Errorf("%+v does not start with %+v", id, testPrefix)
	}
}

func TestImplementsIIdentity(t *testing.T) {
	// Given
	testId := NewIdentity(TEST_PREFIX)
	// When
	b := ImplementsIIdentity(testId)
	// Then
	assert.True(t, b)
}
