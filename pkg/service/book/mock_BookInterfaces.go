// Code generated by mockery v2.16.0. DO NOT EDIT.

package bookservice

import (
	mock "github.com/stretchr/testify/mock"
	bookrepository "github.com/wachayathorn/golang-basic-structure/pkg/repository/book"

	model "github.com/wachayathorn/golang-basic-structure/pkg/model"
)

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

// AddBook provides a mock function with given fields: book
func (_m *MockBookInterfaces) AddBook(book model.Book) (bookrepository.Book, error) {
	ret := _m.Called(book)

	var r0 bookrepository.Book
	if rf, ok := ret.Get(0).(func(model.Book) bookrepository.Book); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Get(0).(bookrepository.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Book) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBookInterfaces_AddBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBook'
type MockBookInterfaces_AddBook_Call struct {
	*mock.Call
}

// AddBook is a helper method to define mock.On call
//   - book model.Book
func (_e *MockBookInterfaces_Expecter) AddBook(book interface{}) *MockBookInterfaces_AddBook_Call {
	return &MockBookInterfaces_AddBook_Call{Call: _e.mock.On("AddBook", book)}
}

func (_c *MockBookInterfaces_AddBook_Call) Run(run func(book model.Book)) *MockBookInterfaces_AddBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.Book))
	})
	return _c
}

func (_c *MockBookInterfaces_AddBook_Call) Return(_a0 bookrepository.Book, _a1 error) *MockBookInterfaces_AddBook_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetBooks provides a mock function with given fields:
func (_m *MockBookInterfaces) GetBooks() []bookrepository.Book {
	ret := _m.Called()

	var r0 []bookrepository.Book
	if rf, ok := ret.Get(0).(func() []bookrepository.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bookrepository.Book)
		}
	}

	return r0
}

// MockBookInterfaces_GetBooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBooks'
type MockBookInterfaces_GetBooks_Call struct {
	*mock.Call
}

// GetBooks is a helper method to define mock.On call
func (_e *MockBookInterfaces_Expecter) GetBooks() *MockBookInterfaces_GetBooks_Call {
	return &MockBookInterfaces_GetBooks_Call{Call: _e.mock.On("GetBooks")}
}

func (_c *MockBookInterfaces_GetBooks_Call) Run(run func()) *MockBookInterfaces_GetBooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBookInterfaces_GetBooks_Call) Return(_a0 []bookrepository.Book) *MockBookInterfaces_GetBooks_Call {
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
