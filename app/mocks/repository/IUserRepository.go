// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	models "github.com/FrancoLiberali/stori_challenge/app/models"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// GetByEmail provides a mock function with given fields: db, email
func (_m *IUserRepository) GetByEmail(db *gorm.DB, email string) (*models.User, error) {
	ret := _m.Called(db, email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) (*models.User, error)); ok {
		return rf(db, email)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) *models.User); ok {
		r0 = rf(db, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string) error); ok {
		r1 = rf(db, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: db, user
func (_m *IUserRepository) Save(db *gorm.DB, user *models.User) error {
	ret := _m.Called(db, user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.User) error); ok {
		r0 = rf(db, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIUserRepository creates a new instance of IUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserRepository {
	mock := &IUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
