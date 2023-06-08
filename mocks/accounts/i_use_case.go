// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	dto "CRM-Services-Task/accounts/dto"
	entity "CRM-Services-Task/accounts/entity"

	mock "github.com/stretchr/testify/mock"
)

// IUseCase is an autogenerated mock type for the IUseCase type
type IUseCase struct {
	mock.Mock
}

type IUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *IUseCase) EXPECT() *IUseCase_Expecter {
	return &IUseCase_Expecter{mock: &_m.Mock}
}

// ActivateAdmin provides a mock function with given fields: a, aar, activeTrue
func (_m *IUseCase) ActivateAdmin(a *entity.Actor, aar *dto.ActivateAdminRequest, activeTrue map[string]interface{}) error {
	ret := _m.Called(a, aar, activeTrue)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.ActivateAdminRequest, map[string]interface{}) error); ok {
		r0 = rf(a, aar, activeTrue)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IUseCase_ActivateAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ActivateAdmin'
type IUseCase_ActivateAdmin_Call struct {
	*mock.Call
}

// ActivateAdmin is a helper method to define mock.On call
//   - a *entity.Actor
//   - aar *dto.ActivateAdminRequest
//   - activeTrue map[string]interface{}
func (_e *IUseCase_Expecter) ActivateAdmin(a interface{}, aar interface{}, activeTrue interface{}) *IUseCase_ActivateAdmin_Call {
	return &IUseCase_ActivateAdmin_Call{Call: _e.mock.On("ActivateAdmin", a, aar, activeTrue)}
}

func (_c *IUseCase_ActivateAdmin_Call) Run(run func(a *entity.Actor, aar *dto.ActivateAdminRequest, activeTrue map[string]interface{})) *IUseCase_ActivateAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.ActivateAdminRequest), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *IUseCase_ActivateAdmin_Call) Return(_a0 error) *IUseCase_ActivateAdmin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IUseCase_ActivateAdmin_Call) RunAndReturn(run func(*entity.Actor, *dto.ActivateAdminRequest, map[string]interface{}) error) *IUseCase_ActivateAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// ApproveAdmin provides a mock function with given fields: a, ar, ra, update, ra2
func (_m *IUseCase) ApproveAdmin(a *entity.Actor, ar *dto.ApproveRequest, ra *entity.RegisterApproval, update map[string]interface{}, ra2 map[string]interface{}) error {
	ret := _m.Called(a, ar, ra, update, ra2)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.ApproveRequest, *entity.RegisterApproval, map[string]interface{}, map[string]interface{}) error); ok {
		r0 = rf(a, ar, ra, update, ra2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IUseCase_ApproveAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApproveAdmin'
type IUseCase_ApproveAdmin_Call struct {
	*mock.Call
}

// ApproveAdmin is a helper method to define mock.On call
//   - a *entity.Actor
//   - ar *dto.ApproveRequest
//   - ra *entity.RegisterApproval
//   - update map[string]interface{}
//   - ra2 map[string]interface{}
func (_e *IUseCase_Expecter) ApproveAdmin(a interface{}, ar interface{}, ra interface{}, update interface{}, ra2 interface{}) *IUseCase_ApproveAdmin_Call {
	return &IUseCase_ApproveAdmin_Call{Call: _e.mock.On("ApproveAdmin", a, ar, ra, update, ra2)}
}

func (_c *IUseCase_ApproveAdmin_Call) Run(run func(a *entity.Actor, ar *dto.ApproveRequest, ra *entity.RegisterApproval, update map[string]interface{}, ra2 map[string]interface{})) *IUseCase_ApproveAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.ApproveRequest), args[2].(*entity.RegisterApproval), args[3].(map[string]interface{}), args[4].(map[string]interface{}))
	})
	return _c
}

func (_c *IUseCase_ApproveAdmin_Call) Return(_a0 error) *IUseCase_ApproveAdmin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IUseCase_ApproveAdmin_Call) RunAndReturn(run func(*entity.Actor, *dto.ApproveRequest, *entity.RegisterApproval, map[string]interface{}, map[string]interface{}) error) *IUseCase_ApproveAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAdmin provides a mock function with given fields: a
func (_m *IUseCase) CreateAdmin(a *entity.Actor) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IUseCase_CreateAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAdmin'
type IUseCase_CreateAdmin_Call struct {
	*mock.Call
}

// CreateAdmin is a helper method to define mock.On call
//   - a *entity.Actor
func (_e *IUseCase_Expecter) CreateAdmin(a interface{}) *IUseCase_CreateAdmin_Call {
	return &IUseCase_CreateAdmin_Call{Call: _e.mock.On("CreateAdmin", a)}
}

func (_c *IUseCase_CreateAdmin_Call) Run(run func(a *entity.Actor)) *IUseCase_CreateAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor))
	})
	return _c
}

func (_c *IUseCase_CreateAdmin_Call) Return(_a0 error) *IUseCase_CreateAdmin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IUseCase_CreateAdmin_Call) RunAndReturn(run func(*entity.Actor) error) *IUseCase_CreateAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// DeactivateAdmin provides a mock function with given fields: a, d, val
func (_m *IUseCase) DeactivateAdmin(a *entity.Actor, d *dto.DeActivateAdminRequest, val map[string]interface{}) error {
	ret := _m.Called(a, d, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.DeActivateAdminRequest, map[string]interface{}) error); ok {
		r0 = rf(a, d, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IUseCase_DeactivateAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeactivateAdmin'
type IUseCase_DeactivateAdmin_Call struct {
	*mock.Call
}

// DeactivateAdmin is a helper method to define mock.On call
//   - a *entity.Actor
//   - d *dto.DeActivateAdminRequest
//   - val map[string]interface{}
func (_e *IUseCase_Expecter) DeactivateAdmin(a interface{}, d interface{}, val interface{}) *IUseCase_DeactivateAdmin_Call {
	return &IUseCase_DeactivateAdmin_Call{Call: _e.mock.On("DeactivateAdmin", a, d, val)}
}

func (_c *IUseCase_DeactivateAdmin_Call) Run(run func(a *entity.Actor, d *dto.DeActivateAdminRequest, val map[string]interface{})) *IUseCase_DeactivateAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.DeActivateAdminRequest), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *IUseCase_DeactivateAdmin_Call) Return(_a0 error) *IUseCase_DeactivateAdmin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IUseCase_DeactivateAdmin_Call) RunAndReturn(run func(*entity.Actor, *dto.DeActivateAdminRequest, map[string]interface{}) error) *IUseCase_DeactivateAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAdmin provides a mock function with given fields: a, d
func (_m *IUseCase) DeleteAdmin(a *entity.Actor, d *dto.DeleteAdminRequest) error {
	ret := _m.Called(a, d)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.DeleteAdminRequest) error); ok {
		r0 = rf(a, d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IUseCase_DeleteAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAdmin'
type IUseCase_DeleteAdmin_Call struct {
	*mock.Call
}

// DeleteAdmin is a helper method to define mock.On call
//   - a *entity.Actor
//   - d *dto.DeleteAdminRequest
func (_e *IUseCase_Expecter) DeleteAdmin(a interface{}, d interface{}) *IUseCase_DeleteAdmin_Call {
	return &IUseCase_DeleteAdmin_Call{Call: _e.mock.On("DeleteAdmin", a, d)}
}

func (_c *IUseCase_DeleteAdmin_Call) Run(run func(a *entity.Actor, d *dto.DeleteAdminRequest)) *IUseCase_DeleteAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.DeleteAdminRequest))
	})
	return _c
}

func (_c *IUseCase_DeleteAdmin_Call) Return(_a0 error) *IUseCase_DeleteAdmin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IUseCase_DeleteAdmin_Call) RunAndReturn(run func(*entity.Actor, *dto.DeleteAdminRequest) error) *IUseCase_DeleteAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// GetActorsByUsername provides a mock function with given fields: username, limit, page
func (_m *IUseCase) GetActorsByUsername(username string, limit string, page string) ([]entity.Actor, error) {
	ret := _m.Called(username, limit, page)

	var r0 []entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) ([]entity.Actor, error)); ok {
		return rf(username, limit, page)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) []entity.Actor); ok {
		r0 = rf(username, limit, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(username, limit, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IUseCase_GetActorsByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActorsByUsername'
type IUseCase_GetActorsByUsername_Call struct {
	*mock.Call
}

// GetActorsByUsername is a helper method to define mock.On call
//   - username string
//   - limit string
//   - page string
func (_e *IUseCase_Expecter) GetActorsByUsername(username interface{}, limit interface{}, page interface{}) *IUseCase_GetActorsByUsername_Call {
	return &IUseCase_GetActorsByUsername_Call{Call: _e.mock.On("GetActorsByUsername", username, limit, page)}
}

func (_c *IUseCase_GetActorsByUsername_Call) Run(run func(username string, limit string, page string)) *IUseCase_GetActorsByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *IUseCase_GetActorsByUsername_Call) Return(_a0 []entity.Actor, _a1 error) *IUseCase_GetActorsByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IUseCase_GetActorsByUsername_Call) RunAndReturn(run func(string, string, string) ([]entity.Actor, error)) *IUseCase_GetActorsByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// GetApprovalList provides a mock function with given fields:
func (_m *IUseCase) GetApprovalList() ([]entity.Actor, error) {
	ret := _m.Called()

	var r0 []entity.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.Actor, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.Actor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Actor)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IUseCase_GetApprovalList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetApprovalList'
type IUseCase_GetApprovalList_Call struct {
	*mock.Call
}

// GetApprovalList is a helper method to define mock.On call
func (_e *IUseCase_Expecter) GetApprovalList() *IUseCase_GetApprovalList_Call {
	return &IUseCase_GetApprovalList_Call{Call: _e.mock.On("GetApprovalList")}
}

func (_c *IUseCase_GetApprovalList_Call) Run(run func()) *IUseCase_GetApprovalList_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IUseCase_GetApprovalList_Call) Return(_a0 []entity.Actor, _a1 error) *IUseCase_GetApprovalList_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IUseCase_GetApprovalList_Call) RunAndReturn(run func() ([]entity.Actor, error)) *IUseCase_GetApprovalList_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: a, lr
func (_m *IUseCase) Login(a *entity.Actor, lr *dto.LoginRequest) error {
	ret := _m.Called(a, lr)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.LoginRequest) error); ok {
		r0 = rf(a, lr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IUseCase_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type IUseCase_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - a *entity.Actor
//   - lr *dto.LoginRequest
func (_e *IUseCase_Expecter) Login(a interface{}, lr interface{}) *IUseCase_Login_Call {
	return &IUseCase_Login_Call{Call: _e.mock.On("Login", a, lr)}
}

func (_c *IUseCase_Login_Call) Run(run func(a *entity.Actor, lr *dto.LoginRequest)) *IUseCase_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.LoginRequest))
	})
	return _c
}

func (_c *IUseCase_Login_Call) Return(_a0 error) *IUseCase_Login_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IUseCase_Login_Call) RunAndReturn(run func(*entity.Actor, *dto.LoginRequest) error) *IUseCase_Login_Call {
	_c.Call.Return(run)
	return _c
}

// RejectAdmin provides a mock function with given fields: a, ar, ra, update, valueRa
func (_m *IUseCase) RejectAdmin(a *entity.Actor, ar *dto.ApproveRequest, ra *entity.RegisterApproval, update map[string]interface{}, valueRa map[string]interface{}) error {
	ret := _m.Called(a, ar, ra, update, valueRa)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.ApproveRequest, *entity.RegisterApproval, map[string]interface{}, map[string]interface{}) error); ok {
		r0 = rf(a, ar, ra, update, valueRa)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IUseCase_RejectAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RejectAdmin'
type IUseCase_RejectAdmin_Call struct {
	*mock.Call
}

// RejectAdmin is a helper method to define mock.On call
//   - a *entity.Actor
//   - ar *dto.ApproveRequest
//   - ra *entity.RegisterApproval
//   - update map[string]interface{}
//   - valueRa map[string]interface{}
func (_e *IUseCase_Expecter) RejectAdmin(a interface{}, ar interface{}, ra interface{}, update interface{}, valueRa interface{}) *IUseCase_RejectAdmin_Call {
	return &IUseCase_RejectAdmin_Call{Call: _e.mock.On("RejectAdmin", a, ar, ra, update, valueRa)}
}

func (_c *IUseCase_RejectAdmin_Call) Run(run func(a *entity.Actor, ar *dto.ApproveRequest, ra *entity.RegisterApproval, update map[string]interface{}, valueRa map[string]interface{})) *IUseCase_RejectAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.ApproveRequest), args[2].(*entity.RegisterApproval), args[3].(map[string]interface{}), args[4].(map[string]interface{}))
	})
	return _c
}

func (_c *IUseCase_RejectAdmin_Call) Return(_a0 error) *IUseCase_RejectAdmin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IUseCase_RejectAdmin_Call) RunAndReturn(run func(*entity.Actor, *dto.ApproveRequest, *entity.RegisterApproval, map[string]interface{}, map[string]interface{}) error) *IUseCase_RejectAdmin_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUseCase creates a new instance of IUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUseCase(t mockConstructorTestingTNewIUseCase) *IUseCase {
	mock := &IUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
