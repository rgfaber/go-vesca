package model

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
)

var (
	BogusConfig       = config.NewConfig()
	BogusGreenhouseID = NewGreenhouseTestID()
	BogusGreenhouse   = NewTestGreenhouse()
	BogusTemp         = 20.0
	BogusHumidity     = 42.0
	BogusSettings     = NewSettings(BogusTemp, BogusHumidity)
	BogusSensor       = NewSensor(BogusConfig.SensorId(), BogusConfig.SensorName())
	BogusFan          = NewFan(BogusConfig.FanId(), BogusConfig.FanName())
	BogusSprinkler    = NewSprinkler(BogusConfig.SprinklerId(), BogusConfig.SprinklerName())
)

func NewTestGreenhouse() *Greenhouse {
	id := BogusGreenhouseID.Value
	name := BogusConfig.GreenhouseName()
	return NewGreenhouse(id, name, BogusSettings, BogusSensor, BogusFan, BogusSprinkler)
}
