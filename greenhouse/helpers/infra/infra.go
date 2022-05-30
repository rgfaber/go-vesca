package infra

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk/infra/memory/store"
)

func LoadGreenhouse(store store.IStore, aggregateId string) *model.Greenhouse {
	res := store.Load(aggregateId)
	if res == nil {
		return nil
	}
	return res.(*model.Greenhouse)
}

func SaveGreenhouse(store store.IStore, state *model.Greenhouse) {
	store.Save(state.ID.Id(), state)
}
