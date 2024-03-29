// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	request "sigma-test/internal/request"

	mock "github.com/stretchr/testify/mock"

	response "sigma-test/internal/response"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

type UserService_Expecter struct {
	mock *mock.Mock
}

func (_m *UserService) EXPECT() *UserService_Expecter {
	return &UserService_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: user
func (_m *UserService) CreateUser(user request.User) (response.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 response.User
	var r1 error
	if rf, ok := ret.Get(0).(func(request.User) (response.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(request.User) response.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(response.User)
	}

	if rf, ok := ret.Get(1).(func(request.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type UserService_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user request.User
func (_e *UserService_Expecter) CreateUser(user interface{}) *UserService_CreateUser_Call {
	return &UserService_CreateUser_Call{Call: _e.mock.On("CreateUser", user)}
}

func (_c *UserService_CreateUser_Call) Run(run func(user request.User)) *UserService_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(request.User))
	})
	return _c
}

func (_c *UserService_CreateUser_Call) Return(_a0 response.User, _a1 error) *UserService_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_CreateUser_Call) RunAndReturn(run func(request.User) (response.User, error)) *UserService_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: id
func (_m *UserService) DeleteUser(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserService_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type UserService_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - id string
func (_e *UserService_Expecter) DeleteUser(id interface{}) *UserService_DeleteUser_Call {
	return &UserService_DeleteUser_Call{Call: _e.mock.On("DeleteUser", id)}
}

func (_c *UserService_DeleteUser_Call) Run(run func(id string)) *UserService_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserService_DeleteUser_Call) Return(_a0 error) *UserService_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserService_DeleteUser_Call) RunAndReturn(run func(string) error) *UserService_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllUsers provides a mock function with given fields:
func (_m *UserService) GetAllUsers() ([]response.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []response.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]response.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []response.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]response.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetAllUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllUsers'
type UserService_GetAllUsers_Call struct {
	*mock.Call
}

// GetAllUsers is a helper method to define mock.On call
func (_e *UserService_Expecter) GetAllUsers() *UserService_GetAllUsers_Call {
	return &UserService_GetAllUsers_Call{Call: _e.mock.On("GetAllUsers")}
}

func (_c *UserService_GetAllUsers_Call) Run(run func()) *UserService_GetAllUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UserService_GetAllUsers_Call) Return(_a0 []response.User, _a1 error) *UserService_GetAllUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetAllUsers_Call) RunAndReturn(run func() ([]response.User, error)) *UserService_GetAllUsers_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *UserService) GetUserByEmail(email string) (response.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 response.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (response.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) response.User); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(response.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type UserService_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - email string
func (_e *UserService_Expecter) GetUserByEmail(email interface{}) *UserService_GetUserByEmail_Call {
	return &UserService_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", email)}
}

func (_c *UserService_GetUserByEmail_Call) Run(run func(email string)) *UserService_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserService_GetUserByEmail_Call) Return(_a0 response.User, _a1 error) *UserService_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetUserByEmail_Call) RunAndReturn(run func(string) (response.User, error)) *UserService_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserById provides a mock function with given fields: id
func (_m *UserService) GetUserById(id string) (response.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserById")
	}

	var r0 response.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (response.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) response.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(response.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_GetUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserById'
type UserService_GetUserById_Call struct {
	*mock.Call
}

// GetUserById is a helper method to define mock.On call
//   - id string
func (_e *UserService_Expecter) GetUserById(id interface{}) *UserService_GetUserById_Call {
	return &UserService_GetUserById_Call{Call: _e.mock.On("GetUserById", id)}
}

func (_c *UserService_GetUserById_Call) Run(run func(id string)) *UserService_GetUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserService_GetUserById_Call) Return(_a0 response.User, _a1 error) *UserService_GetUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_GetUserById_Call) RunAndReturn(run func(string) (response.User, error)) *UserService_GetUserById_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: body
func (_m *UserService) Login(body request.User) (string, error) {
	ret := _m.Called(body)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(request.User) (string, error)); ok {
		return rf(body)
	}
	if rf, ok := ret.Get(0).(func(request.User) string); ok {
		r0 = rf(body)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(request.User) error); ok {
		r1 = rf(body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type UserService_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - body request.User
func (_e *UserService_Expecter) Login(body interface{}) *UserService_Login_Call {
	return &UserService_Login_Call{Call: _e.mock.On("Login", body)}
}

func (_c *UserService_Login_Call) Run(run func(body request.User)) *UserService_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(request.User))
	})
	return _c
}

func (_c *UserService_Login_Call) Return(_a0 string, _a1 error) *UserService_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_Login_Call) RunAndReturn(run func(request.User) (string, error)) *UserService_Login_Call {
	_c.Call.Return(run)
	return _c
}

// SignUp provides a mock function with given fields: body
func (_m *UserService) SignUp(body request.User) (response.User, error) {
	ret := _m.Called(body)

	if len(ret) == 0 {
		panic("no return value specified for SignUp")
	}

	var r0 response.User
	var r1 error
	if rf, ok := ret.Get(0).(func(request.User) (response.User, error)); ok {
		return rf(body)
	}
	if rf, ok := ret.Get(0).(func(request.User) response.User); ok {
		r0 = rf(body)
	} else {
		r0 = ret.Get(0).(response.User)
	}

	if rf, ok := ret.Get(1).(func(request.User) error); ok {
		r1 = rf(body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_SignUp_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SignUp'
type UserService_SignUp_Call struct {
	*mock.Call
}

// SignUp is a helper method to define mock.On call
//   - body request.User
func (_e *UserService_Expecter) SignUp(body interface{}) *UserService_SignUp_Call {
	return &UserService_SignUp_Call{Call: _e.mock.On("SignUp", body)}
}

func (_c *UserService_SignUp_Call) Run(run func(body request.User)) *UserService_SignUp_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(request.User))
	})
	return _c
}

func (_c *UserService_SignUp_Call) Return(_a0 response.User, _a1 error) *UserService_SignUp_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_SignUp_Call) RunAndReturn(run func(request.User) (response.User, error)) *UserService_SignUp_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: id, user
func (_m *UserService) UpdateUser(id string, user request.User) (response.User, error) {
	ret := _m.Called(id, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 response.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, request.User) (response.User, error)); ok {
		return rf(id, user)
	}
	if rf, ok := ret.Get(0).(func(string, request.User) response.User); ok {
		r0 = rf(id, user)
	} else {
		r0 = ret.Get(0).(response.User)
	}

	if rf, ok := ret.Get(1).(func(string, request.User) error); ok {
		r1 = rf(id, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserService_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type UserService_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - id string
//   - user request.User
func (_e *UserService_Expecter) UpdateUser(id interface{}, user interface{}) *UserService_UpdateUser_Call {
	return &UserService_UpdateUser_Call{Call: _e.mock.On("UpdateUser", id, user)}
}

func (_c *UserService_UpdateUser_Call) Run(run func(id string, user request.User)) *UserService_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(request.User))
	})
	return _c
}

func (_c *UserService_UpdateUser_Call) Return(_a0 response.User, _a1 error) *UserService_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserService_UpdateUser_Call) RunAndReturn(run func(string, request.User) (response.User, error)) *UserService_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
