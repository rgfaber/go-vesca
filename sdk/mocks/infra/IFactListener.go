// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// IFactListener is an autogenerated mock type for the IFactListener type
type IFactListener struct {
	mock.Mock
}

// Activate provides a mock function with given fields:
func (_m *IFactListener) Activate() {
	_m.Called()
}

// NewIFactListener creates a new instance of IFactListener. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIFactListener(t testing.TB) *IFactListener {
	mock := &IFactListener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
