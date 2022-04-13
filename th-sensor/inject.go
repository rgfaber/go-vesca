package main

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/features/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

var (
	Store    = infra.NewStore()
	Root     = model.NewRoot(config.Cfg)
	MemBus   = dec.NewDECBus()
	Features = []domain.IFeature{
		initialize.NewFeature(Root, MemBus),
		//kill.NewFeature(Root, MemBus),
		//measure.NewFeature(Root, MemBus),
	}
)
