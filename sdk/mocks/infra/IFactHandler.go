// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// IFactHandler is an autogenerated mock type for the IFactHandler type
type IFactHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: data
func (_m *IFactHandler) Handle(data []byte) {
	_m.Called(data)
}

// NewIFactHandler creates a new instance of IFactHandler. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIFactHandler(t testing.TB) *IFactHandler {
	mock := &IFactHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
