// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "gitlab.com/alfred_soegiarto/training-clean-arch/domain"
)

// ItemRepository is an autogenerated mock type for the ItemRepository type
type ItemRepository struct {
	mock.Mock
}

// FetchAll provides a mock function with given fields: user_id
func (_m *ItemRepository) FetchAll(user_id int64) ([]domain.Item, error) {
	ret := _m.Called(user_id)

	var r0 []domain.Item
	if rf, ok := ret.Get(0).(func(int64) []domain.Item); ok {
		r0 = rf(user_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: user_id, appName
func (_m *ItemRepository) GetByName(user_id int64, appName string) (domain.Item, error) {
	ret := _m.Called(user_id, appName)

	var r0 domain.Item
	if rf, ok := ret.Get(0).(func(int64, string) domain.Item); ok {
		r0 = rf(user_id, appName)
	} else {
		r0 = ret.Get(0).(domain.Item)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(user_id, appName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
