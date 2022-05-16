package interfaces

type IDECBus interface {
	Subscribe(topic string, fn interface{}) error
	SubscribeOnce(topic string, fn interface{}) error
	HasCallback(topic string) bool
	Unsubscribe(topic string, fn interface{}) error
	Publish(topic string, msg ...interface{})
	SubscribeAsync(topic string, fn interface{}, transactional bool) interface{}
	SubscribeOnceAsync(topic string, fn interface{}) error
	WaitAsync()
}

func ImplementsIDECBus(bus IDECBus) bool {
	return true
}
