// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	dto "CRM-Services-Task/customers/dto"
	entity "CRM-Services-Task/customers/entity"

	mock "github.com/stretchr/testify/mock"
)

// ICustomer is an autogenerated mock type for the ICustomer type
type ICustomer struct {
	mock.Mock
}

type ICustomer_Expecter struct {
	mock *mock.Mock
}

func (_m *ICustomer) EXPECT() *ICustomer_Expecter {
	return &ICustomer_Expecter{mock: &_m.Mock}
}

// CreateCustomer provides a mock function with given fields: c
func (_m *ICustomer) CreateCustomer(c *entity.Customer) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Customer) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ICustomer_CreateCustomer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomer'
type ICustomer_CreateCustomer_Call struct {
	*mock.Call
}

// CreateCustomer is a helper method to define mock.On call
//   - c *entity.Customer
func (_e *ICustomer_Expecter) CreateCustomer(c interface{}) *ICustomer_CreateCustomer_Call {
	return &ICustomer_CreateCustomer_Call{Call: _e.mock.On("CreateCustomer", c)}
}

func (_c *ICustomer_CreateCustomer_Call) Run(run func(c *entity.Customer)) *ICustomer_CreateCustomer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Customer))
	})
	return _c
}

func (_c *ICustomer_CreateCustomer_Call) Return(_a0 error) *ICustomer_CreateCustomer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ICustomer_CreateCustomer_Call) RunAndReturn(run func(*entity.Customer) error) *ICustomer_CreateCustomer_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteCustomer provides a mock function with given fields: c, d
func (_m *ICustomer) DeleteCustomer(c *entity.Customer, d *dto.DeleteCustomerRequest) error {
	ret := _m.Called(c, d)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Customer, *dto.DeleteCustomerRequest) error); ok {
		r0 = rf(c, d)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ICustomer_DeleteCustomer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteCustomer'
type ICustomer_DeleteCustomer_Call struct {
	*mock.Call
}

// DeleteCustomer is a helper method to define mock.On call
//   - c *entity.Customer
//   - d *dto.DeleteCustomerRequest
func (_e *ICustomer_Expecter) DeleteCustomer(c interface{}, d interface{}) *ICustomer_DeleteCustomer_Call {
	return &ICustomer_DeleteCustomer_Call{Call: _e.mock.On("DeleteCustomer", c, d)}
}

func (_c *ICustomer_DeleteCustomer_Call) Run(run func(c *entity.Customer, d *dto.DeleteCustomerRequest)) *ICustomer_DeleteCustomer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Customer), args[1].(*dto.DeleteCustomerRequest))
	})
	return _c
}

func (_c *ICustomer_DeleteCustomer_Call) Return(_a0 error) *ICustomer_DeleteCustomer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ICustomer_DeleteCustomer_Call) RunAndReturn(run func(*entity.Customer, *dto.DeleteCustomerRequest) error) *ICustomer_DeleteCustomer_Call {
	_c.Call.Return(run)
	return _c
}

// GetCustomersByNameAndEmail provides a mock function with given fields: firstname, email, limit, page
func (_m *ICustomer) GetCustomersByNameAndEmail(firstname string, email string, limit string, page string) ([]entity.Customer, error) {
	ret := _m.Called(firstname, email, limit, page)

	var r0 []entity.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) ([]entity.Customer, error)); ok {
		return rf(firstname, email, limit, page)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string) []entity.Customer); ok {
		r0 = rf(firstname, email, limit, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(firstname, email, limit, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ICustomer_GetCustomersByNameAndEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCustomersByNameAndEmail'
type ICustomer_GetCustomersByNameAndEmail_Call struct {
	*mock.Call
}

// GetCustomersByNameAndEmail is a helper method to define mock.On call
//   - firstname string
//   - email string
//   - limit string
//   - page string
func (_e *ICustomer_Expecter) GetCustomersByNameAndEmail(firstname interface{}, email interface{}, limit interface{}, page interface{}) *ICustomer_GetCustomersByNameAndEmail_Call {
	return &ICustomer_GetCustomersByNameAndEmail_Call{Call: _e.mock.On("GetCustomersByNameAndEmail", firstname, email, limit, page)}
}

func (_c *ICustomer_GetCustomersByNameAndEmail_Call) Run(run func(firstname string, email string, limit string, page string)) *ICustomer_GetCustomersByNameAndEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *ICustomer_GetCustomersByNameAndEmail_Call) Return(_a0 []entity.Customer, _a1 error) *ICustomer_GetCustomersByNameAndEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ICustomer_GetCustomersByNameAndEmail_Call) RunAndReturn(run func(string, string, string, string) ([]entity.Customer, error)) *ICustomer_GetCustomersByNameAndEmail_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewICustomer interface {
	mock.TestingT
	Cleanup(func())
}

// NewICustomer creates a new instance of ICustomer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICustomer(t mockConstructorTestingTNewICustomer) *ICustomer {
	mock := &ICustomer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}