// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	cart "ecommerceapi/features/cart"

	mock "github.com/stretchr/testify/mock"
)

// CartService is an autogenerated mock type for the CartService type
type CartService struct {
	mock.Mock
}

// AddCart provides a mock function with given fields: token, productId, newCart
func (_m *CartService) AddCart(token interface{}, productId uint, newCart cart.Core) (cart.Core, error) {
	ret := _m.Called(token, productId, newCart)

	var r0 cart.Core
	if rf, ok := ret.Get(0).(func(interface{}, uint, cart.Core) cart.Core); ok {
		r0 = rf(token, productId, newCart)
	} else {
		r0 = ret.Get(0).(cart.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint, cart.Core) error); ok {
		r1 = rf(token, productId, newCart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCart provides a mock function with given fields: token
func (_m *CartService) DeleteCart(token interface{}) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShowCart provides a mock function with given fields: token
func (_m *CartService) ShowCart(token interface{}) ([]cart.Core, error) {
	ret := _m.Called(token)

	var r0 []cart.Core
	if rf, ok := ret.Get(0).(func(interface{}) []cart.Core); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cart.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCart provides a mock function with given fields: token, cartId, updCart
func (_m *CartService) UpdateCart(token interface{}, cartId uint, updCart cart.Core) (cart.Core, error) {
	ret := _m.Called(token, cartId, updCart)

	var r0 cart.Core
	if rf, ok := ret.Get(0).(func(interface{}, uint, cart.Core) cart.Core); ok {
		r0 = rf(token, cartId, updCart)
	} else {
		r0 = ret.Get(0).(cart.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, uint, cart.Core) error); ok {
		r1 = rf(token, cartId, updCart)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCartService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartService creates a new instance of CartService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartService(t mockConstructorTestingTNewCartService) *CartService {
	mock := &CartService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
