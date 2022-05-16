// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/rgfaber/go-vesca/sdk/interfaces"
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	time "time"
)

// INatsBus is an autogenerated mock type for the INatsBus type
type INatsBus struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *INatsBus) Close() {
	_m.Called()
}

// Listen provides a mock function with given fields: topic, handler
func (_m *INatsBus) Listen(topic string, handler interfaces.IFactHandler) {
	_m.Called(topic, handler)
}

// Publish provides a mock function with given fields: topic, data
func (_m *INatsBus) Publish(topic string, data []byte) {
	_m.Called(topic, data)
}

// Request provides a mock function with given fields: topic, data, timeout
func (_m *INatsBus) Request(topic string, data []byte, timeout time.Duration) []byte {
	ret := _m.Called(topic, data, timeout)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, []byte, time.Duration) []byte); ok {
		r0 = rf(topic, data, timeout)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// Respond provides a mock function with given fields: topic, handler
func (_m *INatsBus) Respond(topic string, handler interfaces.IHopeHandler) {
	_m.Called(topic, handler)
}

// NewINatsBus creates a new instance of INatsBus. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewINatsBus(t testing.TB) *INatsBus {
	mock := &INatsBus{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
