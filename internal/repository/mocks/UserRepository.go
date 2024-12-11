// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	entity "hub_management_service/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *UserRepository) Create(user *entity.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: id
func (_m *UserRepository) FindByID(id uint) (*entity.User, error) {
	ret := _m.Called(id)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*entity.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *entity.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByTeamID provides a mock function with given fields: teamID
func (_m *UserRepository) FindUserByTeamID(teamID uint) ([]entity.User, error) {
	ret := _m.Called(teamID)

	var r0 []entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]entity.User, error)); ok {
		return rf(teamID)
	}
	if rf, ok := ret.Get(0).(func(uint) []entity.User); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}