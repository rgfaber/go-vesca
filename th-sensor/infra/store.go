package infra

import (
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Store struct {
	Items map[string]*model.Root
}

func (s *Store) Load(id string) *model.Root {
	return s.Items[id]
}

func (s *Store) Save(item model.Root) {
	s.Items[item.SensorId.Id()] = &item
}

func NewStore() domain.IStore {
	return &Store{
		Items: make(map[string]*model.Root),
	}
}
