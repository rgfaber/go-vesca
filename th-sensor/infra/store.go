package infra

import (
	"github.com/rgfaber/go-vesca/sdk/dec"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Store struct {
	Items map[string]*model.Root
}

func (s *Store) Load(id string) interface{} {
	return s.Items[id]
}

func (s *Store) Save(item interface{}) {
	i := item.(model.Root)
	s.Items[i.SensorId.Id()] = &i
}

func NewStore() dec.IStore {
	return &Store{
		Items: make(map[string]*model.Root),
	}
}
