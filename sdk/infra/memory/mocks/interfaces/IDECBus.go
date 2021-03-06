// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// IDECBus is an autogenerated mock type for the IDECBus type
type IDECBus struct {
	mock.Mock
}

// HasCallback provides a mock function with given fields: topic
func (_m *IDECBus) HasCallback(topic string) bool {
	ret := _m.Called(topic)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(topic)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Publish provides a mock function with given fields: topic, msg
func (_m *IDECBus) Publish(topic string, msg ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, topic)
	_ca = append(_ca, msg...)
	_m.Called(_ca...)
}

// Subscribe provides a mock function with given fields: topic, fn
func (_m *IDECBus) Subscribe(topic string, fn interface{}) error {
	ret := _m.Called(topic, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeAsync provides a mock function with given fields: topic, fn, transactional
func (_m *IDECBus) SubscribeAsync(topic string, fn interface{}, transactional bool) interface{} {
	ret := _m.Called(topic, fn, transactional)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, interface{}, bool) interface{}); ok {
		r0 = rf(topic, fn, transactional)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// SubscribeOnce provides a mock function with given fields: topic, fn
func (_m *IDECBus) SubscribeOnce(topic string, fn interface{}) error {
	ret := _m.Called(topic, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeOnceAsync provides a mock function with given fields: topic, fn
func (_m *IDECBus) SubscribeOnceAsync(topic string, fn interface{}) error {
	ret := _m.Called(topic, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: topic, fn
func (_m *IDECBus) Unsubscribe(topic string, fn interface{}) error {
	ret := _m.Called(topic, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(topic, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitAsync provides a mock function with given fields:
func (_m *IDECBus) WaitAsync() {
	_m.Called()
}

// NewIDECBus creates a new instance of IDECBus. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDECBus(t testing.TB) *IDECBus {
	mock := &IDECBus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
