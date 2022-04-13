package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAggregate(t *testing.T) {
	bus := dec.NewDECBus()
	store := infra.NewStore()
	id := sdk.NewIdentity(config.GO_VESCA_TH_SENSOR_PREFIX)
	a := NewAggregate(id, store, bus)
	assert.NotNil(t, a)
	assert.Nil(t, a.state)
}

func TestAggregate_Execute(t *testing.T) {
	bus := dec.NewDECBus()
	store := infra.NewStore()
	id := sdk.NewIdentity(config.GO_VESCA_TH_SENSOR_PREFIX)
	a := NewAggregate(id, store, bus)
	cmd := NewCmd()
	rsp, e := a.Execute(cmd)
}
