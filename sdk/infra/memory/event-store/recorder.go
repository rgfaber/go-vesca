package event_store

import "github.com/rgfaber/go-vesca/sdk/domain"

type IRecorder interface {
	Record(evt domain.IEvt)
}
