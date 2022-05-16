package interfaces

type IFactEmitter interface {
	Activate()
}

func ImplementsIFactEmitter(emitter IFactEmitter) bool {
	return true
}
