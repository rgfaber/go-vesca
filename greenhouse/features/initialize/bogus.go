package initialize

import (
	bogus_model "github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core/test"
)

var (
	BogusCmd = NewCmd(bogus_model.BogusGreenhouseID,
		test.TEST_TRACE_ID,
		bogus_model.NewDetails(bogus_model.BogusConfig.GreenhouseName(), ""),
		bogus_model.BogusSettings,
		bogus_model.BogusSensor,
		bogus_model.BogusFan,
		bogus_model.BogusSprinkler)
	BogusEvt = NewEvt(bogus_model.BogusGreenhouseID,
		test.TEST_TRACE_ID,
		bogus_model.BogusGreenhouse)
)
