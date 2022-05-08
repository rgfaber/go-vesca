package interfaces

type IFactEmitter interface {
	Emit(fact IFact)
}

func ImplementsIFactEmitter(emitter IFactEmitter) bool {
	return true
}
