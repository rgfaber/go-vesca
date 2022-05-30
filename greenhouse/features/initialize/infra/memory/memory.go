package memory

import (
	memory_mediator "github.com/rgfaber/go-vesca/sdk/infra/memory/mediator"
	memory_store "github.com/rgfaber/go-vesca/sdk/infra/memory/store"
)

var (
	Store    = memory_store.SingletonStore()
	Mediator = memory_mediator.SingletonDECBus()
)
