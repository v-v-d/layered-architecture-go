// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	cart "layered-arch/internal/domain/cart"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// CartsRepository is an autogenerated mock type for the CartsRepository type
type CartsRepository struct {
	mock.Mock
}

// Clear provides a mock function with given fields: aCart
func (_m *CartsRepository) Clear(aCart cart.Cart) error {
	ret := _m.Called(aCart)

	if len(ret) == 0 {
		panic("no return value specified for Clear")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(cart.Cart) error); ok {
		r0 = rf(aCart)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: aCart
func (_m *CartsRepository) Create(aCart cart.Cart) (cart.Cart, error) {
	ret := _m.Called(aCart)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 cart.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(cart.Cart) (cart.Cart, error)); ok {
		return rf(aCart)
	}
	if rf, ok := ret.Get(0).(func(cart.Cart) cart.Cart); ok {
		r0 = rf(aCart)
	} else {
		r0 = ret.Get(0).(cart.Cart)
	}

	if rf, ok := ret.Get(1).(func(cart.Cart) error); ok {
		r1 = rf(aCart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Retrieve provides a mock function with given fields: cartId
func (_m *CartsRepository) Retrieve(cartId uuid.UUID) (cart.Cart, error) {
	ret := _m.Called(cartId)

	if len(ret) == 0 {
		panic("no return value specified for Retrieve")
	}

	var r0 cart.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (cart.Cart, error)); ok {
		return rf(cartId)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) cart.Cart); ok {
		r0 = rf(cartId)
	} else {
		r0 = ret.Get(0).(cart.Cart)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(cartId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: aCart
func (_m *CartsRepository) Update(aCart cart.Cart) (cart.Cart, error) {
	ret := _m.Called(aCart)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 cart.Cart
	var r1 error
	if rf, ok := ret.Get(0).(func(cart.Cart) (cart.Cart, error)); ok {
		return rf(aCart)
	}
	if rf, ok := ret.Get(0).(func(cart.Cart) cart.Cart); ok {
		r0 = rf(aCart)
	} else {
		r0 = ret.Get(0).(cart.Cart)
	}

	if rf, ok := ret.Get(1).(func(cart.Cart) error); ok {
		r1 = rf(aCart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCartsRepository creates a new instance of CartsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCartsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CartsRepository {
	mock := &CartsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
