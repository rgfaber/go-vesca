package sdk

import (
	"strings"
	"testing"
)

func TestNewIdentity(t *testing.T) {
	id := NewIdentity("my")
	if id == nil {
		t.Errorf("id was not created")
	}
	if id.Prefix != "my" {
		t.Errorf("Identity has wrong prefix. Expected 'my' got %+v", id.Prefix)
	}
}

func TestIdentity_Id(t *testing.T) {
	testPrefix := "my"
	id := NewIdentity(testPrefix).Id()
	if !strings.HasPrefix(id, testPrefix) {
		t.Errorf("%+v does not start with %+v", id, testPrefix)
	}
	//	parts :=

}

func TestCleanUuid(t *testing.T) {
	uid, _ := cleanUuid()
	if strings.Contains(uid, "-") {
		t.Errorf("%+v contains '-'", uid)
	}
	if len(uid) != 32 {
		t.Errorf("%+v is not 32 characters long", uid)
	}
}

func TestNullUuid(t *testing.T) {
	uid := cleanNullUuid()
	if strings.Contains(uid, "-") {
		t.Errorf("%+v contains '-'", uid)
	}
	if len(uid) != 32 {
		t.Errorf("%+v is not 32 characters long", uid)
	}
}

func TestNewUuid(t *testing.T) {
	uid, _ := newUuid()
	if !strings.Contains(uid, "-") {
		t.Errorf("%+v  does not contain '-'", uid)
	}
	if len(uid) != 36 {
		t.Errorf("%+v is not 36 characters long. Length is: %+v", uid, len(uid))
	}

}
