package main

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/domain/kill"
	"github.com/rgfaber/go-vesca/th-sensor/domain/measure"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

var (
	Root = model.NewRoot(config.Cfg)
	//	Mediator = sdk.NewMediator
	Features = []domain.IFeature{
		initialize.NewFeature(Root, initialize.ChCmds, initialize.ChEvts),
		kill.NewFeature(Root, kill.ChCmds, kill.ChEvts),
		measure.NewFeature(Root, measure.ChCmds, measure.ChEvts),
	}
)
