package bogus

import (
	"github.com/rgfaber/go-vesca/greenhouse/config"
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

var (
	BogusConfig       = config.NewConfig()
	BogusGreenhouseID = model.NewGreenhouseTestID()
	BogusGreenhouse   = NewTestGreenhouse()
	BogusTemp         = 20.0
	BogusHumidity     = 42.0
	BogusSettings     = model.NewSettings(BogusTemp, BogusHumidity)
	BogusSensor       = model.NewSensor(BogusConfig.SensorId(), BogusConfig.SensorName())
	BogusFan          = model.NewFan(BogusConfig.FanId(), BogusConfig.FanName())
	BogusSprinkler    = model.NewSprinkler(BogusConfig.SprinklerId(), BogusConfig.SprinklerName())
	BogusDetails      = model.NewDetails(BogusConfig.GreenhouseName(), "")
)

func NewTestGreenhouse() *model.Greenhouse {
	id := BogusGreenhouseID.Value
	name := BogusConfig.GreenhouseName()
	return model.NewGreenhouse(id, name, BogusSettings, BogusSensor, BogusFan, BogusSprinkler)
}
