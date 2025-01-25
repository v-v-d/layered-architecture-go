// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	customer "layered-arch/internal/domain/customer"

	mock "github.com/stretchr/testify/mock"
)

// AuthSystem is an autogenerated mock type for the AuthSystem type
type AuthSystem struct {
	mock.Mock
}

// GetCustomer provides a mock function with given fields: data
func (_m *AuthSystem) GetCustomer(data string) (customer.Customer, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomer")
	}

	var r0 customer.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (customer.Customer, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(string) customer.Customer); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(customer.Customer)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthSystem creates a new instance of AuthSystem. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthSystem(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthSystem {
	mock := &AuthSystem{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
