package main

import (
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/domain/kill"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

var (
	Cfg                = config.NewConfig()
	Root               = model.NewRoot(Cfg)
	InitializeCommands = initialize.NewCommands(10)
	InitializeEvents   = initialize.NewEvents(10)
	Features           = []domain.IFeature{
		initialize.NewFeature(Root, InitializeCommands, InitializeEvents),
		kill.NewFeature(Root),
	}
)
