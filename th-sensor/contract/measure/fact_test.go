package measure

import "testing"

func TestNewFact(t *testing.T) {
	f := NewFact()
	if f == nil {
		t.Errorf("Could not create kill.Fact!")
	}
}
