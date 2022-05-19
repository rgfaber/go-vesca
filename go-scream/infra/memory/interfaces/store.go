package interfaces

type IStore interface {
	Load(id string) interface{}
	Save(id string, model interface{})
}

func ImplemementsIStore(store IStore) bool {
	return true
}
