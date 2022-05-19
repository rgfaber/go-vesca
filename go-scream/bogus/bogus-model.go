package bogus

import (
	"github.com/rgfaber/go-vesca/sdk/core"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/interfaces"
)

type BogusStatus int

const (
	Created     BogusStatus = 0 << iota
	Initialized BogusStatus = 1
	Started     BogusStatus = 2
	Ended       BogusStatus = 4
)

const (
	BOGUS_PREFIX     = "bogus"
	BOGUS_FACT_TOPIC = "bogus.started"
	BOGUS_HOPE_TOPIC = "bogus.start"
)

type BogusModel struct {
	ID     *core.Identity
	Status BogusStatus
}

func NewBogusModel(id string) BogusModel {
	return BogusModel{
		ID:     core.NewIdentityFrom(BOGUS_PREFIX, id),
		Status: Created,
	}
}

func SaveBogusState(store interfaces.IStore, model BogusModel) {
	store.Save(model.ID.Id(), model)
}

func LoadBogusState(store interfaces.IStore, id string) BogusModel {
	return store.Load(id).(BogusModel)
}
