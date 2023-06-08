// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	dto "CRM-Services-Task/accounts/dto"
	entity "CRM-Services-Task/accounts/entity"

	mock "github.com/stretchr/testify/mock"
)

// IRepo is an autogenerated mock type for the IRepo type
type IRepo struct {
	mock.Mock
}

type IRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *IRepo) EXPECT() *IRepo_Expecter {
	return &IRepo_Expecter{mock: &_m.Mock}
}

// ActivateAdmin provides a mock function with given fields: a, aar, activeTrue
func (_m *IRepo) ActivateAdmin(a *entity.Actor, aar *dto.ActivateAdminRequest, activeTrue map[string]interface{}) error {
	ret := _m.Called(a, aar, activeTrue)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.ActivateAdminRequest, map[string]interface{}) error); ok {
		r0 = rf(a, aar, activeTrue)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IRepo_ActivateAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ActivateAdmin'
type IRepo_ActivateAdmin_Call struct {
	*mock.Call
}

// ActivateAdmin is a helper method to define mock.On call
//   - a *entity.Actor
//   - aar *dto.ActivateAdminRequest
//   - activeTrue map[string]interface{}
func (_e *IRepo_Expecter) ActivateAdmin(a interface{}, aar interface{}, activeTrue interface{}) *IRepo_ActivateAdmin_Call {
	return &IRepo_ActivateAdmin_Call{Call: _e.mock.On("ActivateAdmin", a, aar, activeTrue)}
}

func (_c *IRepo_ActivateAdmin_Call) Run(run func(a *entity.Actor, aar *dto.ActivateAdminRequest, activeTrue map[string]interface{})) *IRepo_ActivateAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.ActivateAdminRequest), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *IRepo_ActivateAdmin_Call) Return(_a0 error) *IRepo_ActivateAdmin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IRepo_ActivateAdmin_Call) RunAndReturn(run func(*entity.Actor, *dto.ActivateAdminRequest, map[string]interface{}) error) *IRepo_ActivateAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// CreateAdmin provides a mock function with given fields: a
func (_m *IRepo) CreateAdmin(a *entity.Actor) error {
	ret := _m.Called(a)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor) error); ok {
		r0 = rf(a)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IRepo_CreateAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAdmin'
type IRepo_CreateAdmin_Call struct {
	*mock.Call
}

// CreateAdmin is a helper method to define mock.On call
//   - a *entity.Actor
func (_e *IRepo_Expecter) CreateAdmin(a interface{}) *IRepo_CreateAdmin_Call {
	return &IRepo_CreateAdmin_Call{Call: _e.mock.On("CreateAdmin", a)}
}

func (_c *IRepo_CreateAdmin_Call) Run(run func(a *entity.Actor)) *IRepo_CreateAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor))
	})
	return _c
}

func (_c *IRepo_CreateAdmin_Call) Return(_a0 error) *IRepo_CreateAdmin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IRepo_CreateAdmin_Call) RunAndReturn(run func(*entity.Actor) error) *IRepo_CreateAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// DeactivateAdmin provides a mock function with given fields: a, d, val
func (_m *IRepo) DeactivateAdmin(a *entity.Actor, d *dto.DeActivateAdminRequest, val map[string]interface{}) error {
	ret := _m.Called(a, d, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.DeActivateAdminRequest, map[string]interface{}) error); ok {
		r0 = rf(a, d, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IRepo_DeactivateAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeactivateAdmin'
type IRepo_DeactivateAdmin_Call struct {
	*mock.Call
}

// DeactivateAdmin is a helper method to define mock.On call
//   - a *entity.Actor
//   - d *dto.DeActivateAdminRequest
//   - val map[string]interface{}
func (_e *IRepo_Expecter) DeactivateAdmin(a interface{}, d interface{}, val interface{}) *IRepo_DeactivateAdmin_Call {
	return &IRepo_DeactivateAdmin_Call{Call: _e.mock.On("DeactivateAdmin", a, d, val)}
}

func (_c *IRepo_DeactivateAdmin_Call) Run(run func(a *entity.Actor, d *dto.DeActivateAdminRequest, val map[string]interface{})) *IRepo_DeactivateAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.DeActivateAdminRequest), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *IRepo_DeactivateAdmin_Call) Return(_a0 error) *IRepo_DeactivateAdmin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IRepo_DeactivateAdmin_Call) RunAndReturn(run func(*entity.Actor, *dto.DeActivateAdminRequest, map[string]interface{}) error) *IRepo_DeactivateAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAdmin provides a mock function with given fields: a, d
func (_m *IRepo) DeleteAdmin(a *entity.Actor, d *dto.DeleteAdminRequest) error {
	ret := _m.Called(a, d)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.DeleteAdminRequest) error); ok {
		r0 = rf(a, d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IRepo_DeleteAdmin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAdmin'
type IRepo_DeleteAdmin_Call struct {
	*mock.Call
}

// DeleteAdmin is a helper method to define mock.On call
//   - a *entity.Actor
//   - d *dto.DeleteAdminRequest
func (_e *IRepo_Expecter) DeleteAdmin(a interface{}, d interface{}) *IRepo_DeleteAdmin_Call {
	return &IRepo_DeleteAdmin_Call{Call: _e.mock.On("DeleteAdmin", a, d)}
}

func (_c *IRepo_DeleteAdmin_Call) Run(run func(a *entity.Actor, d *dto.DeleteAdminRequest)) *IRepo_DeleteAdmin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.DeleteAdminRequest))
	})
	return _c
}

func (_c *IRepo_DeleteAdmin_Call) Return(_a0 error) *IRepo_DeleteAdmin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IRepo_DeleteAdmin_Call) RunAndReturn(run func(*entity.Actor, *dto.DeleteAdminRequest) error) *IRepo_DeleteAdmin_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllActors provides a mock function with given fields: username, limit, page
func (_m *IRepo) FindAllActors(username string, limit string, page string) ([]entity.Actor, error) {
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

// IRepo_FindAllActors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllActors'
type IRepo_FindAllActors_Call struct {
	*mock.Call
}

// FindAllActors is a helper method to define mock.On call
//   - username string
//   - limit string
//   - page string
func (_e *IRepo_Expecter) FindAllActors(username interface{}, limit interface{}, page interface{}) *IRepo_FindAllActors_Call {
	return &IRepo_FindAllActors_Call{Call: _e.mock.On("FindAllActors", username, limit, page)}
}

func (_c *IRepo_FindAllActors_Call) Run(run func(username string, limit string, page string)) *IRepo_FindAllActors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *IRepo_FindAllActors_Call) Return(_a0 []entity.Actor, _a1 error) *IRepo_FindAllActors_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IRepo_FindAllActors_Call) RunAndReturn(run func(string, string, string) ([]entity.Actor, error)) *IRepo_FindAllActors_Call {
	_c.Call.Return(run)
	return _c
}

// FindAllApproval provides a mock function with given fields:
func (_m *IRepo) FindAllApproval() ([]entity.Actor, error) {
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

// IRepo_FindAllApproval_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAllApproval'
type IRepo_FindAllApproval_Call struct {
	*mock.Call
}

// FindAllApproval is a helper method to define mock.On call
func (_e *IRepo_Expecter) FindAllApproval() *IRepo_FindAllApproval_Call {
	return &IRepo_FindAllApproval_Call{Call: _e.mock.On("FindAllApproval")}
}

func (_c *IRepo_FindAllApproval_Call) Run(run func()) *IRepo_FindAllApproval_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IRepo_FindAllApproval_Call) Return(_a0 []entity.Actor, _a1 error) *IRepo_FindAllApproval_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IRepo_FindAllApproval_Call) RunAndReturn(run func() ([]entity.Actor, error)) *IRepo_FindAllApproval_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: a, lr
func (_m *IRepo) Login(a *entity.Actor, lr *dto.LoginRequest) error {
	ret := _m.Called(a, lr)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.LoginRequest) error); ok {
		r0 = rf(a, lr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IRepo_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type IRepo_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - a *entity.Actor
//   - lr *dto.LoginRequest
func (_e *IRepo_Expecter) Login(a interface{}, lr interface{}) *IRepo_Login_Call {
	return &IRepo_Login_Call{Call: _e.mock.On("Login", a, lr)}
}

func (_c *IRepo_Login_Call) Run(run func(a *entity.Actor, lr *dto.LoginRequest)) *IRepo_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.LoginRequest))
	})
	return _c
}

func (_c *IRepo_Login_Call) Return(_a0 error) *IRepo_Login_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IRepo_Login_Call) RunAndReturn(run func(*entity.Actor, *dto.LoginRequest) error) *IRepo_Login_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateActor provides a mock function with given fields: a, ar, update
func (_m *IRepo) UpdateActor(a *entity.Actor, ar *dto.ApproveRequest, update map[string]interface{}) error {
	ret := _m.Called(a, ar, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Actor, *dto.ApproveRequest, map[string]interface{}) error); ok {
		r0 = rf(a, ar, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IRepo_UpdateActor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateActor'
type IRepo_UpdateActor_Call struct {
	*mock.Call
}

// UpdateActor is a helper method to define mock.On call
//   - a *entity.Actor
//   - ar *dto.ApproveRequest
//   - update map[string]interface{}
func (_e *IRepo_Expecter) UpdateActor(a interface{}, ar interface{}, update interface{}) *IRepo_UpdateActor_Call {
	return &IRepo_UpdateActor_Call{Call: _e.mock.On("UpdateActor", a, ar, update)}
}

func (_c *IRepo_UpdateActor_Call) Run(run func(a *entity.Actor, ar *dto.ApproveRequest, update map[string]interface{})) *IRepo_UpdateActor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Actor), args[1].(*dto.ApproveRequest), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *IRepo_UpdateActor_Call) Return(_a0 error) *IRepo_UpdateActor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IRepo_UpdateActor_Call) RunAndReturn(run func(*entity.Actor, *dto.ApproveRequest, map[string]interface{}) error) *IRepo_UpdateActor_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateStatusRA provides a mock function with given fields: ra, ar, val
func (_m *IRepo) UpdateStatusRA(ra *entity.RegisterApproval, ar *dto.ApproveRequest, val map[string]interface{}) error {
	ret := _m.Called(ra, ar, val)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.RegisterApproval, *dto.ApproveRequest, map[string]interface{}) error); ok {
		r0 = rf(ra, ar, val)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IRepo_UpdateStatusRA_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateStatusRA'
type IRepo_UpdateStatusRA_Call struct {
	*mock.Call
}

// UpdateStatusRA is a helper method to define mock.On call
//   - ra *entity.RegisterApproval
//   - ar *dto.ApproveRequest
//   - val map[string]interface{}
func (_e *IRepo_Expecter) UpdateStatusRA(ra interface{}, ar interface{}, val interface{}) *IRepo_UpdateStatusRA_Call {
	return &IRepo_UpdateStatusRA_Call{Call: _e.mock.On("UpdateStatusRA", ra, ar, val)}
}

func (_c *IRepo_UpdateStatusRA_Call) Run(run func(ra *entity.RegisterApproval, ar *dto.ApproveRequest, val map[string]interface{})) *IRepo_UpdateStatusRA_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.RegisterApproval), args[1].(*dto.ApproveRequest), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *IRepo_UpdateStatusRA_Call) Return(_a0 error) *IRepo_UpdateStatusRA_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IRepo_UpdateStatusRA_Call) RunAndReturn(run func(*entity.RegisterApproval, *dto.ApproveRequest, map[string]interface{}) error) *IRepo_UpdateStatusRA_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewIRepo creates a new instance of IRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIRepo(t mockConstructorTestingTNewIRepo) *IRepo {
	mock := &IRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}