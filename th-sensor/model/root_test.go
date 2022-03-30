package model

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"testing"
)

func TestNewRoot(t *testing.T) {
	cfg := config.NewConfig()
	r := NewRoot(cfg)
	if r == nil {
		t.Errorf("No Root was created")
	}
	if &r.ID == nil {
		t.Errorf("Root has no Identity")
	}
}
