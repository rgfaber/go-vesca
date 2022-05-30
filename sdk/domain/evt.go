package domain

import "github.com/rgfaber/go-vesca/sdk/core"

type IEvt interface {
	AggregateId() core.IIdentity
	TraceId() string
	Order() int
}

func ImplementsIEvt(evt IEvt) bool {
	return true
}

type EvtBase struct {
	aggregateId string
	traceId     string
	eventId     string
	version     int
	order       int
	eventType   string
	payload     []byte
	metadata    []byte
}
