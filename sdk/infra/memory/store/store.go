package store

import (
	"sync"
)

type IStore interface {
	Load(id string) interface{}
	Save(id string, model interface{})
}

func ImplemementsIStore(store IStore) bool {
	return true
}

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

func SingletonStore() IStore {
	if singleStore == nil {
		lock.Lock()
		defer lock.Unlock()
		singleStore = &Store{
			Items: make(map[string]interface{}),
		}
	}
	return singleStore
}

func TransientStore() IStore {
	return &Store{
		Items: make(map[string]interface{}),
	}
}
