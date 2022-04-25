package initialize

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	bus := dec.NewDECBus()
	store := infra.NewStore()
	a := NewAggregate(store, bus)
	assert.NotNil(t, a)
	assert.Nil(t, a.state)
}

func TestAggregate_Attempt(t *testing.T) {
	bus := dec.NewDECBus()
	store := infra.NewStore()
	a := NewAggregate(store, bus)
	cfg := config.NewConfig()
	cmd := NewCmd(cfg.SensorId(), cfg.SensorName(), cfg.GreenhouseId(), TEST_TRACE_ID, 15.0, 50.0)
	rsp, e := a.Attempt(cmd)
}
