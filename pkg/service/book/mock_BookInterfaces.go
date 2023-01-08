// Code generated by mockery v2.16.0. DO NOT EDIT.

package book

import mock "github.com/stretchr/testify/mock"

// MockBookInterfaces is an autogenerated mock type for the BookInterfaces type
type MockBookInterfaces struct {
	mock.Mock
}

type MockBookInterfaces_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBookInterfaces) EXPECT() *MockBookInterfaces_Expecter {
	return &MockBookInterfaces_Expecter{mock: &_m.Mock}
}

// AddBook provides a mock function with given fields:
func (_m *MockBookInterfaces) AddBook() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBookInterfaces_AddBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBook'
type MockBookInterfaces_AddBook_Call struct {
	*mock.Call
}

// AddBook is a helper method to define mock.On call
func (_e *MockBookInterfaces_Expecter) AddBook() *MockBookInterfaces_AddBook_Call {
	return &MockBookInterfaces_AddBook_Call{Call: _e.mock.On("AddBook")}
}

func (_c *MockBookInterfaces_AddBook_Call) Run(run func()) *MockBookInterfaces_AddBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBookInterfaces_AddBook_Call) Return(_a0 error) *MockBookInterfaces_AddBook_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockBookInterfaces interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBookInterfaces creates a new instance of MockBookInterfaces. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBookInterfaces(t mockConstructorTestingTNewMockBookInterfaces) *MockBookInterfaces {
	mock := &MockBookInterfaces{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
