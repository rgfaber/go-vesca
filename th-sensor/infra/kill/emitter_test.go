package kill

import "testing"

func TestNewEmitter(t *testing.T) {
	e := NewEmitter()
	if e == nil {
		t.Errorf("Could not create Emitter")
	}
}
