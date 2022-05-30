package infra

import (
	"github.com/rgfaber/go-vesca/sdk/contract"
	"github.com/rgfaber/go-vesca/sdk/domain"
)

type IHopeRequester interface {
	Request(contract.IHope) (domain.IFbk, error)
}

func ImplementsIRequester(requester IHopeRequester) bool {
	return true
}
