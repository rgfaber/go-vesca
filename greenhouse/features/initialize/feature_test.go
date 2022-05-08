package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/config/envars"
	domain2 "github.com/rgfaber/go-vesca/greenhouse/features/initialize/domain"
	initialize_infra_nats "github.com/rgfaber/go-vesca/greenhouse/features/initialize/infra/nats"
	"github.com/rgfaber/go-vesca/greenhouse/features/measure/infra/nats"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
	nats2 "github.com/rgfaber/go-vesca/sdk/nats"
	sdk_nats_config "github.com/rgfaber/go-vesca/sdk/nats/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	Config         = config.NewConfig()
	NatsConfig     = sdk_nats_config.NewNatsConfig()
	TestStore      = sdk.NewStore()
	TestBus        = sdk.NewDECBus()
	TestNats       = nats2.NewNatsBus(NatsConfig)
	TestSerializer = initialize_infra_nats.NewFactSerializer()
	TestEmitter    = nats.NewEmitter(TestBus, TestNats)
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
	f := NewFeature(TestBus, TestStore, TestEmitter)
	// Then
	assert.NotNil(t, f, "Feature domain.Initialize could not be created.")
}

func TestFeature_RespondToInitializeCmd(t *testing.T) {
	// Given
	setEnv()
	f := NewFeature(TestBus, TestStore, TestEmitter)
	cfg := config.NewConfig()
	sensorId := cfg.SensorId()
	greenhouseName := cfg.GreenhouseName()
	greenhouseId := cfg.GreenhouseId()
	sensorName := cfg.SensorName()
	fanId := cfg.FanId()
	fanName := cfg.FanName()
	sprinklerId := cfg.SprinklerId()
	sprinklerName := cfg.SprinklerName()
	temp := 15.0
	hum := 20.6
	traceId := sdk.TEST_TRACE_ID
	receivedEvent := false
	aggregateId := model.NewGreenhouseID(greenhouseId)
	settings := model.NewSettings(temp, hum)
	sensor := model.NewSensor(sensorId, sensorName)
	fan := model.NewFan(fanId, fanName)
	sprinkler := model.NewSprinkler(sprinklerId, sprinklerName)
	cmd := domain2.NewCmd(aggregateId, traceId, greenhouseName, settings, sensor, fan, sprinkler)
	// When
	f.HandleCommand()
	// And
	TestBus.Subscribe(domain2.EVT_TOPIC, func(evt sdk.IEvt) {
		e := evt.(*domain2.Evt)
		receivedEvent = e != nil
	})
	// And

	TestBus.Publish(domain2.CMD_TOPIC, cmd)
	// Then
	assert.True(t, receivedEvent)
}

func TestImplementsIFeature(t *testing.T) {
	// Given
	f := NewFeature(TestBus, TestStore, TestEmitter)
	// When
	imp := sdk.ImplemmentsIFeature(f)
	// Then
	assert.True(t, imp)
}
