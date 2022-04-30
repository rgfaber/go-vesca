package initialize

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/envars"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	TestStore = infra.NewStore()
	TestBus   = dec.NewDECBus()
)

func setEnv() {
	err := os.Setenv(envars.GO_VESCA_TH_SENSOR_ID, sdk.TEST_UUID)

	if err != nil {
		panic(err)
	}
}

func TestNewFeature(t *testing.T) {
	// Given
	setEnv()
	// When
	f := NewFeature(TestBus, TestStore)
	// Then
	assert.NotNil(t, f, "Feature domain.Initialize could not be created.")
}

func TestFeature_Respond(t *testing.T) {
	// Given
	setEnv()
	f := NewFeature(TestBus, TestStore)
	cfg := config.NewConfig()
	sensorId := cfg.SensorId()
	sensorName := cfg.SensorName()
	greenhouseId := cfg.GreenhouseId()
	temp := 15.0
	hum := 20.6
	traceId := sdk.TEST_TRACE_ID
	receivedEvent := false
	// When
	f.Respond()
	// And
	TestBus.Subscribe(initialize.EVT_TOPIC, func(evt dec.IEvt) {
		e := evt.(*initialize.Evt)
		receivedEvent = e != nil
	})
	// And
	TestBus.Publish(initialize.CMD_TOPIC, initialize.NewCmd(sensorId, sensorName, greenhouseId, traceId, temp, hum))
	// Then
	assert.True(t, receivedEvent)
}
