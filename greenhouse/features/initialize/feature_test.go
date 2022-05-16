package initialize

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/config/envars"
	"github.com/rgfaber/go-vesca/greenhouse/model"
	testing2 "github.com/rgfaber/go-vesca/sdk/core/test"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	"github.com/rgfaber/go-vesca/sdk/interfaces"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setEnv() {
	err := os.Setenv(envars.GO_VESCA_TH_SENSOR_ID, testing2.TEST_UUID)
	if err != nil {
		panic(err)
	}
}

func TestNewFeature(t *testing.T) {
	// Given
	setEnv()

	// When
	f := Feature
	// Then
	assert.NotNil(t, f, "Feature domain.Initialize could not be created.")
}

func TestFeature_RespondToInitializeCmd(t *testing.T) {
	// Given
	setEnv()
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
	traceId := testing2.TEST_TRACE_ID
	receivedEvent := false
	aggregateId := model.NewGreenhouseID(greenhouseId)
	settings := model.NewSettings(temp, hum)
	sensor := model.NewSensor(sensorId, sensorName)
	fan := model.NewFan(fanId, fanName)
	sprinkler := model.NewSprinkler(sprinklerId, sprinklerName)
	details := model.NewDetails(greenhouseName, "")
	cmd := NewCmd(aggregateId, traceId, details, settings, sensor, fan, sprinkler)
	aggregate := NewAggregate(Store, Mediator)
	// When
	mediator.SingletonDECBus().Subscribe(EVT_TOPIC, func(evt interfaces.IEvt) {
		e := evt.(*Evt)
		receivedEvent = e != nil
	})
	// And
	fbk := aggregate.Attempt(cmd)
	// then
	assert.NotNil(t, fbk)
	assert.True(t, receivedEvent)
}

func TestImplementsIFeature(t *testing.T) {
	// Given
	f := Feature
	// When
	imp := interfaces.ImplementsIFeature(f)
	// Then
	assert.True(t, imp)
}
