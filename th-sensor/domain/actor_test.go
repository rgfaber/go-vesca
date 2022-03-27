package domain

import (
	"github.com/rgfaber/go-vesca/sdk"
	"testing"
)

func TestNewActor(t *testing.T) {
	act := NewActor(sdk.TEST_UUID)
	if act == nil {
		t.Errorf("Actor could not be created")
	}
	if act.Id == nil {
		t.Errorf("Actor has no Identity")
	}
	if act.Id.Prefix != TH_SENSOR
}
