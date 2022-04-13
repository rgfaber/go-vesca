package model

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRoot(t *testing.T) {
	cfg := config.NewConfig()
	r := NewRoot(cfg)
	assert.NotNil(t, r)
	assert.NotNil(t, r.ID)
}
