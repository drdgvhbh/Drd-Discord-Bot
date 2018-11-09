// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import api "drdgvhbh/discordbot/internal/user/api"
import entity "drdgvhbh/discordbot/internal/user/entity"
import mock "github.com/stretchr/testify/mock"

// UserMapper is an autogenerated mock type for the UserMapper type
type UserMapper struct {
	mock.Mock
}

// MapFrom provides a mock function with given fields: _a0
func (_m *UserMapper) MapFrom(_a0 *entity.User) *api.User {
	ret := _m.Called(_a0)

	var r0 *api.User
	if rf, ok := ret.Get(0).(func(*entity.User) *api.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.User)
		}
	}

	return r0
}

// MapTo provides a mock function with given fields: _a0
func (_m *UserMapper) MapTo(_a0 *api.User) *entity.User {
	ret := _m.Called(_a0)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(*api.User) *entity.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	return r0
}