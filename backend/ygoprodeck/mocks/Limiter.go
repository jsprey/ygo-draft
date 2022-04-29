// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// Limiter is an autogenerated mock type for the Limiter type
type Limiter struct {
	mock.Mock
}

// Wait provides a mock function with given fields: ctx
func (_m *Limiter) Wait(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewLimiter creates a new instance of Limiter. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewLimiter(t testing.TB) *Limiter {
	mock := &Limiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
