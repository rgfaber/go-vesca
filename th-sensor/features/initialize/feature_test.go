package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/envars"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewFeature(t *testing.T) {
	err := os.Setenv(envars.GO_VESCA_TH_SENSOR_ID, sdk.TEST_UUID)
	assert.Nil(t, err)
	cfg := config.NewConfig()
	state := model.NewRoot(cfg)
	assert.NotNil(t, state, "model.Root could not be created for", cfg.SensorId())
	f := NewFeature(state, dec.NewDECBus())
	assert.NotNil(t, f, "Feature domain.Initialize could not be created.")
}

func TestFeature_Respond(t *testing.T) {
	err := os.Setenv(envars.GO_VESCA_TH_SENSOR_ID, sdk.TEST_UUID)
	assert.Nil(t, err)
	cfg := config.NewConfig()
	state := model.NewRoot(cfg)
	assert.NotNil(t, state, "model.Root could not be created for", cfg.SensorId())
	f := NewFeature(state, dec.NewDECBus())
	assert.NotNil(t, f, "Feature domain.Initialize could not be created.")
	err = f.bus.Subscribe(initialize.CMD_TOPIC, func(cmd *initialize.Cmd) {
		assert.NotNil(t, cmd)
		assert.Equal(t, 15.0, cmd.measurement.Temperature)
		assert.Equal(t, 20.0, cmd.measurement.Humidity)
	})
	assert.Nil(t, err)
	f.bus.Publish(initialize.CMD_TOPIC, initialize.NewCmd(15, 20))
}

func TestThatTheCallbackCannotMutateTheMessageContent(t *testing.T) {
	err := os.Setenv(envars.GO_VESCA_TH_SENSOR_ID, sdk.TEST_UUID)
	assert.Nil(t, err)
	cfg := config.NewConfig()
	state := model.NewRoot(cfg)
	assert.NotNil(t, state, "model.Root could not be created for", cfg.SensorId())
	f := NewFeature(state, dec.NewDECBus())
	assert.NotNil(t, f, "Feature domain.Initialize could not be created.")
	err = f.bus.Subscribe(initialize.CMD_TOPIC, func(c initialize.Cmd) {
		assert.NotNil(t, c)
		assert.Equal(t, 15.0, c.measurement.Temperature)
		assert.Equal(t, 20.0, c.measurement.Humidity)
		c.measurement.Temperature = 22.0
		c.measurement.Humidity = 30.0
	})
	assert.Nil(t, err)
	cmd := initialize.NewCmd(15, 20)
	f.bus.Publish(initialize.CMD_TOPIC, *cmd)
	assert.NotEqual(t, 22.0, cmd.measurement.Temperature)
	assert.NotEqual(t, 30.0, cmd.measurement.Humidity)
}
