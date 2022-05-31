package domain

import (
	bogus "github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/core/mocks"
)

var (
	BogusCmd = NewCmd(bogus.BogusGreenhouseID,
		mocks.TEST_TRACE_ID,
		bogus.BogusFan)
	BogusEvt = NewEvt(bogus.BogusGreenhouseID,
		mocks.TEST_TRACE_ID,
		bogus.BogusFan)
)
