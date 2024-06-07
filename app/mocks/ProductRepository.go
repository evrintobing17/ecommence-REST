// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	models "github.com/evrintobing17/ecommence-REST/app/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *ProductRepository) GetAll() (*[]models.Product, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 *[]models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func() (*[]models.Product, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *[]models.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *ProductRepository) GetByID(id int) (*models.Product, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Product, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySellerID provides a mock function with given fields: id
func (_m *ProductRepository) GetBySellerID(id int) (*[]models.Product, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetBySellerID")
	}

	var r0 *[]models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*[]models.Product, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *[]models.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0
func (_m *ProductRepository) Insert(_a0 *models.Product) (*models.Product, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 *models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Product) (*models.Product, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*models.Product) *models.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(*models.Product) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}