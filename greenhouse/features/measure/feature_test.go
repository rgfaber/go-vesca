package measure

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/domain/measure"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFeature(t *testing.T) {
	cfg := config.NewConfig()
	state := model.NewGreenhouse(cfg.SensorId(), cfg.SensorName(), cfg.GreenhouseId())
	ft := NewFeature(state, sdk.NewDECBus())
	assert.NotNil(t, ft)
}

func TestMeasure(t *testing.T) {

	cfg := config.NewConfig()
	state := model.NewGreenhouse(cfg.SensorId(), cfg.SensorName(), cfg.GreenhouseId())
	ft := NewFeature(state, sdk.NewDECBus())
	assert.NotNil(t, ft)
	go ft.Respond()
	ft.bus.Publish(measure.CMD_TOPIC, measure.NewCmd())
}
