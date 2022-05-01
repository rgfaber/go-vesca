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
	if id.prefix != TEST_PREFIX {
		t.Errorf("Identity has wrong prefix. Expected 'my' got %+v", id.prefix)
	}
}

func TestNewDefaultIdentity(t *testing.T) {
	id := NewIdentity("")
	assert.NotNil(t, id)
	assert.Equal(t, "id", id.prefix)
}

func TestNewIdentityFrom(t *testing.T) {
	id := NewIdentityFrom(TEST_PREFIX, TEST_UUID)
	if id == nil {
		t.Errorf("No Identity was created")
	}
	if id.prefix != TEST_PREFIX {
		t.Errorf("Identity has an incorrect prefix. Expected [%+v] Got [%+v]", TEST_PREFIX, id.prefix)
	}
	seed := strings.Replace(TEST_UUID, "-", "", -1)
	if id.value != seed {
		t.Errorf("Identity has an incorrect value. Expected [%+v] Got [%+v]", seed, id.value)
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
