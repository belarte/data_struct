// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Swapper is an autogenerated mock type for the Swapper type
type Swapper struct {
	mock.Mock
}

// Count provides a mock function with given fields:
func (_m *Swapper) Count() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Swap provides a mock function with given fields: left, right
func (_m *Swapper) Swap(left *int, right *int) {
	_m.Called(left, right)
}
