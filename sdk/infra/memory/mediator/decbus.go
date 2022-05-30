package mediator

import (
	"github.com/asaskevich/EventBus"
	"sync"
)

var lock = &sync.Mutex{}

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

type DECBus struct {
	bus EventBus.Bus
}

var singleDecBus *DECBus

func SingletonDECBus() IDECBus {
	if singleDecBus == nil {
		lock.Lock()
		defer lock.Unlock()
		singleDecBus = &DECBus{bus: EventBus.New()}
	}
	return singleDecBus
}

func TransientDECBus() IDECBus {
	return &DECBus{bus: EventBus.New()}
}

func (b *DECBus) Subscribe(topic string, fn interface{}) error {
	return b.bus.Subscribe(topic, fn)
}

func (b *DECBus) SubscribeOnce(topic string, fn interface{}) error {
	return b.bus.SubscribeOnce(topic, fn)
}

func (b *DECBus) HasCallback(topic string) bool {
	return b.bus.HasCallback(topic)
}

func (b *DECBus) Unsubscribe(topic string, fn interface{}) error {
	return b.bus.Unsubscribe(topic, fn)
}

func (b *DECBus) Publish(topic string, msg ...interface{}) {
	b.bus.Publish(topic, msg...)
}

func (b *DECBus) SubscribeAsync(topic string, fn interface{}, transactional bool) interface{} {
	return b.bus.SubscribeAsync(topic, fn, transactional)
}

func (b *DECBus) SubscribeOnceAsync(topic string, fn interface{}) error {
	return b.bus.SubscribeOnceAsync(topic, fn)
}

func (b *DECBus) WaitAsync() {
	b.bus.WaitAsync()
}
