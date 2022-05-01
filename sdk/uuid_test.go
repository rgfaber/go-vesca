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

func TestCleanUuid(t *testing.T) {
	uid, _ := CleanUuid()
	if strings.Contains(uid, "-") {
		t.Errorf("%+v contains '-'", uid)
	}
	if len(uid) != 32 {
		t.Errorf("%+v is not 32 characters long", uid)
	}
}

func TestNullUuid(t *testing.T) {
	uid := CleanNullId()
	if strings.Contains(uid, "-") {
		t.Errorf("%+v contains '-'", uid)
	}
	if len(uid) != 32 {
		t.Errorf("%+v is not 32 characters long", uid)
	}
}

func TestNewUuid(t *testing.T) {
	uid, _ := NewUuid()
	if !strings.Contains(uid, "-") {
		t.Errorf("%+v  does not contain '-'", uid)
	}
	if len(uid) != 36 {
		t.Errorf("%+v is not 36 characters long. Length is: %+v", uid, len(uid))
	}
}
