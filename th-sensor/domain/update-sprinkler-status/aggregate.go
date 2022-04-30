package update_sprinkler_status

import "github.com/rgfaber/go-vesca/sdk/dec"

type Aggregate struct {
	store dec.IStore
	bus   dec.IDECBus
}
