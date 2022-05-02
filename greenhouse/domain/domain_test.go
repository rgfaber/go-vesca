package domain

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
	"github.com/rgfaber/go-vesca/th-sensor/model"
	"testing"
)

func TestSaveState(t *testing.T) {
	// Given
	cfg := config.NewConfig()
	sensorId := cfg.SensorId()
	sensorName := cfg.SensorName()
	greenhouseId := cfg.GreenhouseId()
	store := infra.NewStore()
	id := model.NewGreenhouseID(sensorId)
	state := model.NewGreenhouse(*id, sensorName, greenhouseId)
	// When
	SaveState(store, state)
}