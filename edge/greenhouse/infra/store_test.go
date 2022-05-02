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
	mIn := model.NewGreenhouse(cfg.SensorId(), cfg.SensorName(), cfg.GreenhouseId())
	mIn.Status = model.Initialized
	s.Save(*mIn)
	mIn.Status = model.Killed
	m := s.Load(mIn.SensorId.Id())
	assert.NotNil(t, m)
	assert.Equal(t, model.Initialized, m.Status)
	m.Status = model.Measuring
	s.Save(*m)
	m = s.Load(m.SensorId.Id())
	assert.NotNil(t, m)
	assert.Equal(t, model.Measuring, m.Status)

}
