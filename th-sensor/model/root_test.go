package model

import "testing"

func TestNewRoot(t *testing.T) {
	r := NewRoot()
	if r == nil {
		t.Errorf("No Root was created")
	}
	if &r.Id == nil {
		t.Errorf("Root has no Identity")
	}
}
