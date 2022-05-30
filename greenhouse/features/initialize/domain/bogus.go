package domain

import (
	bogus_model "github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/greenhouse/model/bogus"
	"github.com/rgfaber/go-vesca/sdk/core/mocks"
)

var (
	BogusCmd = NewCmd(bogus.BogusGreenhouseID,
		mocks.TEST_TRACE_ID,
		bogus_model.NewDetails(bogus.BogusConfig.GreenhouseName(), ""),
		bogus.BogusSettings,
		bogus.BogusSensor,
		bogus.BogusFan,
		bogus.BogusSprinkler)
	BogusEvt = NewEvt(bogus.BogusGreenhouseID,
		mocks.TEST_TRACE_ID,
		bogus.BogusGreenhouse)
)
