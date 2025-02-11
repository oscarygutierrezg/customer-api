// Code generated by mockery v2.46.2. DO NOT EDIT.

package ports

import (
	domain "customer-api/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockClientRepository is an autogenerated mock type for the MockClientRepository type
type MockClientRepository struct {
	mock.Mock
}

// GetClientByID provides a mock function with given fields: id
func (_m *MockClientRepository) GetClientByID(id string) (*domain.Client, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetClientByID")
	}

	var r0 *domain.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Client, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Client); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockClientRepository creates a new instance of MockClientRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClientRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClientRepository {
	mock := &MockClientRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
