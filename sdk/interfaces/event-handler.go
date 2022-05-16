package interfaces

type IEventHandler interface {
	Activate()
}

func ImplementsIEventHandler(handler IEventHandler) bool {
	return true
}

func BuildHandlers(handlers ...IEventHandler) []IEventHandler {
	return handlers
}
