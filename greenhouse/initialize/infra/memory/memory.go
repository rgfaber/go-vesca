package memory

import (
	memory_mediator "github.com/rgfaber/go-vesca/go-scream/infra/memory/mediator"
	memory_store "github.com/rgfaber/go-vesca/go-scream/infra/memory/store"
)

var (
	Store    = memory_store.SingletonStore()
	Mediator = memory_mediator.SingletonDECBus()
)
