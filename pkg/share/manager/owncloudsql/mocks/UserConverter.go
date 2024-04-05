// Copyright 2018-2022 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

// Code generated by mockery v2.40.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
)

// UserConverter is an autogenerated mock type for the UserConverter type
type UserConverter struct {
	mock.Mock
}

type UserConverter_Expecter struct {
	mock *mock.Mock
}

func (_m *UserConverter) EXPECT() *UserConverter_Expecter {
	return &UserConverter_Expecter{mock: &_m.Mock}
}

// GetUser provides a mock function with given fields: userid
func (_m *UserConverter) GetUser(userid *userv1beta1.UserId) (*userv1beta1.User, error) {
	ret := _m.Called(userid)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 *userv1beta1.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*userv1beta1.UserId) (*userv1beta1.User, error)); ok {
		return rf(userid)
	}
	if rf, ok := ret.Get(0).(func(*userv1beta1.UserId) *userv1beta1.User); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*userv1beta1.UserId) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserConverter_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type UserConverter_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - userid *userv1beta1.UserId
func (_e *UserConverter_Expecter) GetUser(userid interface{}) *UserConverter_GetUser_Call {
	return &UserConverter_GetUser_Call{Call: _e.mock.On("GetUser", userid)}
}

func (_c *UserConverter_GetUser_Call) Run(run func(userid *userv1beta1.UserId)) *UserConverter_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*userv1beta1.UserId))
	})
	return _c
}

func (_c *UserConverter_GetUser_Call) Return(_a0 *userv1beta1.User, _a1 error) *UserConverter_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserConverter_GetUser_Call) RunAndReturn(run func(*userv1beta1.UserId) (*userv1beta1.User, error)) *UserConverter_GetUser_Call {
	_c.Call.Return(run)
	return _c
}

// UserIDToUserName provides a mock function with given fields: ctx, userid
func (_m *UserConverter) UserIDToUserName(ctx context.Context, userid *userv1beta1.UserId) (string, error) {
	ret := _m.Called(ctx, userid)

	if len(ret) == 0 {
		panic("no return value specified for UserIDToUserName")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *userv1beta1.UserId) (string, error)); ok {
		return rf(ctx, userid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *userv1beta1.UserId) string); ok {
		r0 = rf(ctx, userid)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *userv1beta1.UserId) error); ok {
		r1 = rf(ctx, userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserConverter_UserIDToUserName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserIDToUserName'
type UserConverter_UserIDToUserName_Call struct {
	*mock.Call
}

// UserIDToUserName is a helper method to define mock.On call
//   - ctx context.Context
//   - userid *userv1beta1.UserId
func (_e *UserConverter_Expecter) UserIDToUserName(ctx interface{}, userid interface{}) *UserConverter_UserIDToUserName_Call {
	return &UserConverter_UserIDToUserName_Call{Call: _e.mock.On("UserIDToUserName", ctx, userid)}
}

func (_c *UserConverter_UserIDToUserName_Call) Run(run func(ctx context.Context, userid *userv1beta1.UserId)) *UserConverter_UserIDToUserName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*userv1beta1.UserId))
	})
	return _c
}

func (_c *UserConverter_UserIDToUserName_Call) Return(_a0 string, _a1 error) *UserConverter_UserIDToUserName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserConverter_UserIDToUserName_Call) RunAndReturn(run func(context.Context, *userv1beta1.UserId) (string, error)) *UserConverter_UserIDToUserName_Call {
	_c.Call.Return(run)
	return _c
}

// UserNameToUserID provides a mock function with given fields: ctx, username
func (_m *UserConverter) UserNameToUserID(ctx context.Context, username string) (*userv1beta1.UserId, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for UserNameToUserID")
	}

	var r0 *userv1beta1.UserId
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*userv1beta1.UserId, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *userv1beta1.UserId); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*userv1beta1.UserId)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserConverter_UserNameToUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserNameToUserID'
type UserConverter_UserNameToUserID_Call struct {
	*mock.Call
}

// UserNameToUserID is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *UserConverter_Expecter) UserNameToUserID(ctx interface{}, username interface{}) *UserConverter_UserNameToUserID_Call {
	return &UserConverter_UserNameToUserID_Call{Call: _e.mock.On("UserNameToUserID", ctx, username)}
}

func (_c *UserConverter_UserNameToUserID_Call) Run(run func(ctx context.Context, username string)) *UserConverter_UserNameToUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserConverter_UserNameToUserID_Call) Return(_a0 *userv1beta1.UserId, _a1 error) *UserConverter_UserNameToUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserConverter_UserNameToUserID_Call) RunAndReturn(run func(context.Context, string) (*userv1beta1.UserId, error)) *UserConverter_UserNameToUserID_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserConverter creates a new instance of UserConverter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserConverter(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserConverter {
	mock := &UserConverter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
