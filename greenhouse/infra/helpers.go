package infra

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
)

func LoadGreenhouse(store sdk.IStore, aggregateId string) *model.Greenhouse {
	return store.Load(aggregateId).(*model.Greenhouse)
}

func SaveGreenhouse(store sdk.IStore, state *model.Greenhouse) {
	store.Save(state.ID.Id(), state)
}
