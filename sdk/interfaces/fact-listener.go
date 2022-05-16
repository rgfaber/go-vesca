package interfaces

type IFactListener interface {
	Activate()
}

func ImplementsIListener(ls IFactListener) bool {
	return true
}

func BuildListeners(listeners ...IFactListener) []IFactListener {
	return listeners
}
