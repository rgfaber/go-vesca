package store

import (
	"github.com/rgfaber/go-vesca/sdk/infra/memory/interfaces"
	"sync"
)

var lock = &sync.Mutex{}

var singleStore *Store

type Store struct {
	Items map[string]interface{}
}

func (s *Store) Load(id string) interface{} {
	return s.Items[id]
}

func (s *Store) Save(id string, item interface{}) {
	s.Items[id] = item
}

func SingletonStore() interfaces.IStore {
	if singleStore == nil {
		lock.Lock()
		defer lock.Unlock()
		singleStore = &Store{
			Items: make(map[string]interface{}),
		}
	}
	return singleStore
}

func TransientStore() interfaces.IStore {
	return &Store{
		Items: make(map[string]interface{}),
	}
}
