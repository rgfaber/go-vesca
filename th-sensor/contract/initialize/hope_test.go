package initialize

import "testing"

func TestNewHope(t *testing.T) {
	h := NewHope()
	if h == nil {
		t.Errorf("Could not create measure.Hope!")
	}
}
