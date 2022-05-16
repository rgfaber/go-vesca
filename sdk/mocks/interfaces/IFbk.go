// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// IFbk is an autogenerated mock type for the IFbk type
type IFbk struct {
	mock.Mock
}

// AggregateId provides a mock function with given fields:
func (_m *IFbk) AggregateId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Error provides a mock function with given fields:
func (_m *IFbk) Error() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsSuccess provides a mock function with given fields:
func (_m *IFbk) IsSuccess() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TraceId provides a mock function with given fields:
func (_m *IFbk) TraceId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewIFbk creates a new instance of IFbk. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIFbk(t testing.TB) *IFbk {
	mock := &IFbk{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
