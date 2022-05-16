package interfaces

type IEventTrigger interface {
	Activate()
}

func ImplementsIEventTrigger(trigger IEventTrigger) bool {
	return true
}
