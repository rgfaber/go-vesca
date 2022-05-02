package domain

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/infra"
	"github.com/rgfaber/go-vesca/greenhouse/model"
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
