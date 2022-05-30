package domain

import "github.com/rgfaber/go-vesca/sdk/core"

type ICmd interface {
	AggregateId() core.IIdentity
	TraceId() string
}

func ImplementsICmd(cmd ICmd) bool {
	return true
}
