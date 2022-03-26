package sdk

import (
	"strings"
	"testing"
)

func TestCleanUuid(t *testing.T) {
	uid := cleanUuid()
	if strings.Contains(uid, "-") {
		t.Errorf("%+v contains '-'", uid)
	}
	if len(uid) != 32 {
		t.Errorf("%+v is not 32 characters long", uid)
	}
}
