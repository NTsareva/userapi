// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	models "userapi/internal/models"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserRepository) CreateUser(user *models.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUsers provides a mock function with given fields:
func (_m *UserRepository) GetUsers() ([]models.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []models.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersBy provides a mock function with given fields: startDate, endDate, minAge, maxAge
func (_m *UserRepository) GetUsersBy(startDate time.Time, endDate time.Time, minAge int, maxAge int) ([]models.User, int64, error) {
	ret := _m.Called(startDate, endDate, minAge, maxAge)

	if len(ret) == 0 {
		panic("no return value specified for GetUsersBy")
	}

	var r0 []models.User
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, int, int) ([]models.User, int64, error)); ok {
		return rf(startDate, endDate, minAge, maxAge)
	}
	if rf, ok := ret.Get(0).(func(time.Time, time.Time, int, int) []models.User); ok {
		r0 = rf(startDate, endDate, minAge, maxAge)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(time.Time, time.Time, int, int) int64); ok {
		r1 = rf(startDate, endDate, minAge, maxAge)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(time.Time, time.Time, int, int) error); ok {
		r2 = rf(startDate, endDate, minAge, maxAge)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
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
