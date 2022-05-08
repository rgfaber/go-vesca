package bogus

import "github.com/rgfaber/go-vesca/sdk"

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
	ID     *sdk.Identity
	Status BogusStatus
}

func NewBogusModel(id string) BogusModel {
	return BogusModel{
		ID:     sdk.NewIdentityFrom(BOGUS_PREFIX, id),
		Status: Created,
	}
}

func SaveBogusState(store sdk.IStore, model BogusModel) {
	store.Save(model.ID.Id(), model)
}

func LoadBogusState(store sdk.IStore, id string) BogusModel {
	return store.Load(id).(BogusModel)
}
