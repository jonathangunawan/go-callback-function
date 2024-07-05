// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UsecaseItf is an autogenerated mock type for the UsecaseItf type
type UsecaseItf struct {
	mock.Mock
}

// SomeService provides a mock function with given fields: data1, data2
func (_m *UsecaseItf) SomeService(data1 int, data2 int) error {
	ret := _m.Called(data1, data2)

	if len(ret) == 0 {
		panic("no return value specified for SomeService")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(data1, data2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUsecaseItf creates a new instance of UsecaseItf. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsecaseItf(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsecaseItf {
	mock := &UsecaseItf{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}