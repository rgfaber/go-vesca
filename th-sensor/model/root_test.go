package model

import "testing"

func TestNewRoot(t *testing.T) {
	r := &Root{}
	if r == nil {
		t.Errorf("No Root was created")
	}
}
