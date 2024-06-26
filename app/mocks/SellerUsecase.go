// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	models "github.com/evrintobing17/ecommence-REST/app/models"
	mock "github.com/stretchr/testify/mock"
)

// SellerUsecase is an autogenerated mock type for the SellerUsecase type
type SellerUsecase struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: userId
func (_m *SellerUsecase) DeleteByID(userId int) error {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: email, password
func (_m *SellerUsecase) Login(email string, password string) (*models.Seller, string, error) {
	ret := _m.Called(email, password)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *models.Seller
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (*models.Seller, string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *models.Seller); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Seller)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RefreshAccessJWT provides a mock function with given fields: userID
func (_m *SellerUsecase) RefreshAccessJWT(userID int) (string, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for RefreshAccessJWT")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (string, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: username, email, password, age, isbuyer
func (_m *SellerUsecase) Register(username string, email string, password string, age string, isbuyer bool) (*models.Seller, error) {
	ret := _m.Called(username, email, password, age, isbuyer)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 *models.Seller
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, bool) (*models.Seller, error)); ok {
		return rf(username, email, password, age, isbuyer)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string, bool) *models.Seller); ok {
		r0 = rf(username, email, password, age, isbuyer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Seller)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string, bool) error); ok {
		r1 = rf(username, email, password, age, isbuyer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: updateData
func (_m *SellerUsecase) Update(updateData map[string]interface{}) (*models.Seller, error) {
	ret := _m.Called(updateData)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *models.Seller
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}) (*models.Seller, error)); ok {
		return rf(updateData)
	}
	if rf, ok := ret.Get(0).(func(map[string]interface{}) *models.Seller); ok {
		r0 = rf(updateData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Seller)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSellerUsecase creates a new instance of SellerUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSellerUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *SellerUsecase {
	mock := &SellerUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
