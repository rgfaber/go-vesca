package infra

import (
	"github.com/rgfaber/go-vesca/go-scream/infra/memory/interfaces"
	"github.com/rgfaber/go-vesca/greenhouse/model"
)

func LoadGreenhouse(store interfaces.IStore, aggregateId string) *model.Greenhouse {
	res := store.Load(aggregateId)
	if res == nil {
		return nil
	}
	return res.(*model.Greenhouse)
}

func SaveGreenhouse(store interfaces.IStore, state *model.Greenhouse) {
	store.Save(state.ID.Id(), state)
}
