package infra

type IFactEmitter interface {
	Activate()
}

func ImplementsIFactEmitter(emitter IFactEmitter) bool {
	return true
}
