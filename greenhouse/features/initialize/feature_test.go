package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/config/envars"
	"github.com/rgfaber/go-vesca/greenhouse/domain/initialize"
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	TestStore = infra.NewStore()
	TestBus   = sdk.NewDECBus()
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

func TestFeature_RespondToInitializeCmd(t *testing.T) {
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
	aggregateId := model.NewGreenhouseID(sensorId)
	cmd := initialize.NewCmd(aggregateId, sensorName, greenhouseId, traceId, temp, hum)
	// When
	f.StartResponding()
	// And
	TestBus.Subscribe(initialize.EVT_TOPIC, func(evt sdk.IEvt) {
		e := evt.(*initialize.Evt)
		receivedEvent = e != nil
	})
	// And

	TestBus.Publish(initialize.CMD_TOPIC, cmd)
	// Then
	assert.True(t, receivedEvent)
}

func TestImplementsIFeature(t *testing.T) {
	// Given
	b := sdk.NewDECBus()
	s := infra.NewStore()
	f := NewFeature(b, s)
	// When
	imp := sdk.ImplemmentsIFeature(f)
	// Then
	assert.True(t, imp)
}
