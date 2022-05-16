package store

import "sync"

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

func NewStore() *Store {
	if singleStore == nil {
		lock.Lock()
		defer lock.Unlock()
		singleStore = &Store{
			Items: make(map[string]interface{}),
		}
	}
	return singleStore
}
