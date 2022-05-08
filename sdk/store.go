package sdk

type Store struct {
	Items map[string]interface{}
}

func (s *Store) Load(id string) interface{} {
	return s.Items[id]
}

func (s *Store) Save(id string, item interface{}) {
	s.Items[id] = item
}

func NewStore() IStore {
	return &Store{
		Items: make(map[string]interface{}),
	}
}

type IStore interface {
	Load(id string) interface{}
	Save(id string, model interface{})
}

func ImplemementsIStore(store IStore) bool {
	return true
}
