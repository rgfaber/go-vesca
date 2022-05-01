package main

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/features/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/infra"
)

var (
	MemStore = infra.NewStore()
	MemBus   = dec.NewDECBus()
	Features = []dec.IFeature{
		initialize.NewFeature(MemBus, MemStore),
		//kill.NewFeature(Root, MemBus),
		//measure.NewFeature(MemBus, MemStore),
	}
)
