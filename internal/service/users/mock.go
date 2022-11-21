// Code generated by mockery v2.14.0. DO NOT EDIT.

package users

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IServiceMock is an autogenerated mock type for the IServiceMock type
type IServiceMock struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, input
func (_m *IServiceMock) CreateUser(ctx context.Context, input CreateUserInput) (UserResponse, error) {
	ret := _m.Called(ctx, input)

	var r0 UserResponse
	if rf, ok := ret.Get(0).(func(context.Context, CreateUserInput) UserResponse); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(UserResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, CreateUserInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx
func (_m *IServiceMock) GetUsers(ctx context.Context) ([]UserEmailResponse, error) {
	ret := _m.Called(ctx)

	var r0 []UserEmailResponse
	if rf, ok := ret.Get(0).(func(context.Context) []UserEmailResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]UserEmailResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIService creates a new instance of IServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIService(t mockConstructorTestingTNewIService) *IServiceMock {
	mock := &IServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}