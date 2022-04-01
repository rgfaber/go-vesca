package main

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/domain/kill"
	"github.com/rgfaber/go-vesca/th-sensor/domain/measure"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

var (
	Root     = model.NewRoot(config.Cfg)
	MemBus   = dec.NewDECBus()
	Features = []domain.IFeature{
		initialize.NewFeature(Root, MemBus),
		kill.NewFeature(Root, MemBus),
		measure.NewFeature(Root, MemBus),
	}
)
