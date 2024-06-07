// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	models "github.com/evrintobing17/ecommence-REST/app/models"
	mock "github.com/stretchr/testify/mock"
)

// OrderUsecase is an autogenerated mock type for the OrderUsecase type
type OrderUsecase struct {
	mock.Mock
}

// AcceptOrder provides a mock function with given fields: sellerID, itemID, address
func (_m *OrderUsecase) AcceptOrder(sellerID int, itemID int, address string) (*models.Order, error) {
	ret := _m.Called(sellerID, itemID, address)

	if len(ret) == 0 {
		panic("no return value specified for AcceptOrder")
	}

	var r0 *models.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string) (*models.Order, error)); ok {
		return rf(sellerID, itemID, address)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) *models.Order); ok {
		r0 = rf(sellerID, itemID, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(sellerID, itemID, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrder provides a mock function with given fields: buyerID, itemID, quantity, address
func (_m *OrderUsecase) CreateOrder(buyerID int, itemID int, quantity int, address string) (*models.Order, error) {
	ret := _m.Called(buyerID, itemID, quantity, address)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 *models.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, int, string) (*models.Order, error)); ok {
		return rf(buyerID, itemID, quantity, address)
	}
	if rf, ok := ret.Get(0).(func(int, int, int, string) *models.Order); ok {
		r0 = rf(buyerID, itemID, quantity, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, int, string) error); ok {
		r1 = rf(buyerID, itemID, quantity, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListOrder provides a mock function with given fields: _a0
func (_m *OrderUsecase) GetListOrder(_a0 int) (*[]models.Order, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetListOrder")
	}

	var r0 *[]models.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*[]models.Order, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *[]models.Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSellerListOrder provides a mock function with given fields: _a0
func (_m *OrderUsecase) GetSellerListOrder(_a0 int) (*[]models.Order, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetSellerListOrder")
	}

	var r0 *[]models.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*[]models.Order, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *[]models.Order); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrderUsecase creates a new instance of OrderUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderUsecase {
	mock := &OrderUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
