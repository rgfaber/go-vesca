package main

import (
	"github.com/rgfaber/go-vesca/sdk"
	"github.com/rgfaber/go-vesca/th-sensor/config"
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/domain/initialize"
	"github.com/rgfaber/go-vesca/th-sensor/domain/kill"
	"github.com/rgfaber/go-vesca/th-sensor/domain/measure"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

var (
	Root     = model.NewRoot(config.Cfg)
	Mediator = sdk.NewMediator(10)
	Features = []domain.IFeature{
		initialize.NewFeature(Root, Mediator),
		kill.NewFeature(Root, Mediator),
		measure.NewFeature(Root, Mediator),
	}
)
