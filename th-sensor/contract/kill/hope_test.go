package kill

import "testing"

func TestNewHope(t *testing.T) {
	h := NewHope()
	if h == nil {
		t.Error("Could not create kill.Hope!")
	}
}
