package interfaces

import (
	"github.com/rgfaber/go-vesca/sdk"
)

type IFeature interface {
	Bus() sdk.IDECBus
	Store() sdk.IStore
	Aggregate() IAggregate
	Run()
	HandleDomainEvents()
	HandleCommand()
}

func ImplemmentsIFeature(feature IFeature) bool {
	return true
}
