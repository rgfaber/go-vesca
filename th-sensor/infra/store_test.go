package infra

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStore(t *testing.T) {
	s := NewStore()
	assert.NotNil(t, s)
}

func TestIfWeCanSaveAndLoadTheModel(t *testing.T) {
	s := NewStore()
	cfg := config.NewConfig()
	mIn := model.NewRoot(cfg)
	mIn.Status = model.Initialized
	s.Save(*mIn)
	m := s.Load()
	assert.NotNil(t, m)
	assert.Equal(t, model.Initialized, m.Status)
}
