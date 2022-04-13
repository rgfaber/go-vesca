package infra

import (
	"github.com/rgfaber/go-vesca/th-sensor/domain"
	"github.com/rgfaber/go-vesca/th-sensor/model"
)

type Store struct {
	Model model.Root
}

func (s *Store) Load() model.Root {
	return s.Model
}

func (s *Store) Save(model model.Root) {
	s.Model = model
}

func NewStore() domain.IStore {
	return &Store{}
}
