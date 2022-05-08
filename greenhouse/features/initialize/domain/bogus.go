package domain

import (
	bogus_model "github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
)

var (
	BogusCmd = NewCmd(bogus_model.BogusGreenhouseID,
		sdk.TEST_TRACE_ID,
		bogus_model.BogusConfig.GreenhouseName(),
		bogus_model.BogusSettings,
		bogus_model.BogusSensor,
		bogus_model.BogusFan,
		bogus_model.BogusSprinkler)
	BogusEvt = NewEvt(bogus_model.BogusGreenhouseID,
		sdk.TEST_TRACE_ID,
		bogus_model.BogusSettings)
)
