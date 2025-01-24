// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	application "layered-arch/internal/application"

	mock "github.com/stretchr/testify/mock"
)

// TrxManager is an autogenerated mock type for the TrxManager type
type TrxManager struct {
	mock.Mock
}

// Carts provides a mock function with no fields
func (_m *TrxManager) Carts() application.CartsRepository {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Carts")
	}

	var r0 application.CartsRepository
	if rf, ok := ret.Get(0).(func() application.CartsRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(application.CartsRepository)
		}
	}

	return r0
}

// Commit provides a mock function with no fields
func (_m *TrxManager) Commit() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Items provides a mock function with no fields
func (_m *TrxManager) Items() application.ItemsRepository {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Items")
	}

	var r0 application.ItemsRepository
	if rf, ok := ret.Get(0).(func() application.ItemsRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(application.ItemsRepository)
		}
	}

	return r0
}

// Rollback provides a mock function with no fields
func (_m *TrxManager) Rollback() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Shutdown provides a mock function with no fields
func (_m *TrxManager) Shutdown() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Shutdown")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewTrxManager creates a new instance of TrxManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTrxManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *TrxManager {
	mock := &TrxManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
