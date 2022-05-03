package infra

import (
	"github.com/rgfaber/go-vesca/greenhouse/model"
	"github.com/rgfaber/go-vesca/sdk"
)

type Store struct {
	Items map[string]*model.Greenhouse
}

func (s *Store) Load(id string) interface{} {
	return s.Items[id]
}

func (s *Store) Save(item interface{}) {
	i := item.(*model.Greenhouse)
	s.Items[i.ID.Id()] = i
}

func NewStore() sdk.IStore {
	return &Store{
		Items: make(map[string]*model.Greenhouse),
	}
}
