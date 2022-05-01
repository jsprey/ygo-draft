// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	model "ygodraft/backend/model"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// YgoClient is an autogenerated mock type for the YgoClient type
type YgoClient struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *YgoClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllCards provides a mock function with given fields:
func (_m *YgoClient) GetAllCards() (*[]*model.Card, error) {
	ret := _m.Called()

	var r0 *[]*model.Card
	if rf, ok := ret.Get(0).(func() *[]*model.Card); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]*model.Card)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCard provides a mock function with given fields: id
func (_m *YgoClient) GetCard(id int) (*model.Card, error) {
	ret := _m.Called(id)

	var r0 *model.Card
	if rf, ok := ret.Get(0).(func(int) *model.Card); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Card)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAllCards provides a mock function with given fields: cards
func (_m *YgoClient) SaveAllCards(cards *[]*model.Card) error {
	ret := _m.Called(cards)

	var r0 error
	if rf, ok := ret.Get(0).(func(*[]*model.Card) error); ok {
		r0 = rf(cards)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveCard provides a mock function with given fields: api
func (_m *YgoClient) SaveCard(card *model.Card) error {
	ret := _m.Called(card)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Card) error); ok {
		r0 = rf(card)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewYgoClient creates a new instance of YgoClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewYgoClient(t testing.TB) *YgoClient {
	mock := &YgoClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
