package interfaces

type IEventHandler interface {
	Activate()
	Subscribe(events chan IEvt, transactional bool)
}

func ImplementsIEventHandler(handler IEventHandler) bool {
	return true
}

func BuildHandlers(handlers ...IEventHandler) []IEventHandler {
	return handlers
}
