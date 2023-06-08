// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IAuth is an autogenerated mock type for the IAuth type
type IAuth struct {
	mock.Mock
}

type IAuth_Expecter struct {
	mock *mock.Mock
}

func (_m *IAuth) EXPECT() *IAuth_Expecter {
	return &IAuth_Expecter{mock: &_m.Mock}
}

// HashPassword provides a mock function with given fields: pass
func (_m *IAuth) HashPassword(pass string) (string, error) {
	ret := _m.Called(pass)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(pass)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(pass)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IAuth_HashPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HashPassword'
type IAuth_HashPassword_Call struct {
	*mock.Call
}

// HashPassword is a helper method to define mock.On call
//   - pass string
func (_e *IAuth_Expecter) HashPassword(pass interface{}) *IAuth_HashPassword_Call {
	return &IAuth_HashPassword_Call{Call: _e.mock.On("HashPassword", pass)}
}

func (_c *IAuth_HashPassword_Call) Run(run func(pass string)) *IAuth_HashPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *IAuth_HashPassword_Call) Return(_a0 string, _a1 error) *IAuth_HashPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IAuth_HashPassword_Call) RunAndReturn(run func(string) (string, error)) *IAuth_HashPassword_Call {
	_c.Call.Return(run)
	return _c
}

// VerifyPassword provides a mock function with given fields: hashedPassword, candidatePassword
func (_m *IAuth) VerifyPassword(hashedPassword string, candidatePassword string) error {
	ret := _m.Called(hashedPassword, candidatePassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(hashedPassword, candidatePassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IAuth_VerifyPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifyPassword'
type IAuth_VerifyPassword_Call struct {
	*mock.Call
}

// VerifyPassword is a helper method to define mock.On call
//   - hashedPassword string
//   - candidatePassword string
func (_e *IAuth_Expecter) VerifyPassword(hashedPassword interface{}, candidatePassword interface{}) *IAuth_VerifyPassword_Call {
	return &IAuth_VerifyPassword_Call{Call: _e.mock.On("VerifyPassword", hashedPassword, candidatePassword)}
}

func (_c *IAuth_VerifyPassword_Call) Run(run func(hashedPassword string, candidatePassword string)) *IAuth_VerifyPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *IAuth_VerifyPassword_Call) Return(_a0 error) *IAuth_VerifyPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuth_VerifyPassword_Call) RunAndReturn(run func(string, string) error) *IAuth_VerifyPassword_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIAuth interface {
	mock.TestingT
	Cleanup(func())
}

// NewIAuth creates a new instance of IAuth. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIAuth(t mockConstructorTestingTNewIAuth) *IAuth {
	mock := &IAuth{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
