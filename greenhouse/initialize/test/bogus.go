package test

import (
	"github.com/rgfaber/go-vesca/go-scream/core/test"
	domain2 "github.com/rgfaber/go-vesca/greenhouse/initialize/domain"
	bogus_model "github.com/rgfaber/go-vesca/greenhouse/model"
)

var (
	BogusCmd = domain2.NewCmd(bogus_model.BogusGreenhouseID,
		test.TEST_TRACE_ID,
		bogus_model.NewDetails(bogus_model.BogusConfig.GreenhouseName(), ""),
		bogus_model.BogusSettings,
		bogus_model.BogusSensor,
		bogus_model.BogusFan,
		bogus_model.BogusSprinkler)
	BogusEvt = domain2.NewEvt(bogus_model.BogusGreenhouseID,
		test.TEST_TRACE_ID,
		bogus_model.BogusGreenhouse)
)
