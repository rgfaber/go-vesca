package initialize

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
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

func TestAggregate_Execute(t *testing.T) {
	bus := dec.NewDECBus()
	store := infra.NewStore()
	a := NewAggregate(store, bus)
	cmd := NewCmd()
	rsp, e := a.Execute(cmd)
}
