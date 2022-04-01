package measure

import (
	"fmt"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFeature(t *testing.T) {
	cfg := config.NewConfig()
	state := model.NewRoot(cfg)
	ft := NewFeature(state, dec.NewDECBus())
	assert.NotNil(t, ft)
}

func TestMeasure(t *testing.T) {
	cfg := config.NewConfig()
	state := model.NewRoot(cfg)
	ft := NewFeature(state, dec.NewDECBus())
	assert.NotNil(t, ft)
	err := ft.bus.Subscribe(CMD_TOPIC, func(cmd *Cmd) {
		fmt.Println("Received:", *cmd)
	})
	if err != nil {
		return
	}
	go ft.measure()
	select {}
}
