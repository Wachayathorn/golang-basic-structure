// Code generated by mockery v2.16.0. DO NOT EDIT.

package client

import (
	mock "github.com/stretchr/testify/mock"
	bookrepository "github.com/wachayathorn/golang-basic-structure/pkg/repository/book"

	model "github.com/wachayathorn/golang-basic-structure/pkg/model"
)

// MockClientAPI is an autogenerated mock type for the ClientAPI type
type MockClientAPI struct {
	mock.Mock
}

type MockClientAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClientAPI) EXPECT() *MockClientAPI_Expecter {
	return &MockClientAPI_Expecter{mock: &_m.Mock}
}

// AddBook provides a mock function with given fields: book
func (_m *MockClientAPI) AddBook(book model.Book) (bookrepository.Book, error) {
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

// MockClientAPI_AddBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBook'
type MockClientAPI_AddBook_Call struct {
	*mock.Call
}

// AddBook is a helper method to define mock.On call
//   - book model.Book
func (_e *MockClientAPI_Expecter) AddBook(book interface{}) *MockClientAPI_AddBook_Call {
	return &MockClientAPI_AddBook_Call{Call: _e.mock.On("AddBook", book)}
}

func (_c *MockClientAPI_AddBook_Call) Run(run func(book model.Book)) *MockClientAPI_AddBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.Book))
	})
	return _c
}

func (_c *MockClientAPI_AddBook_Call) Return(_a0 bookrepository.Book, _a1 error) *MockClientAPI_AddBook_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetBooks provides a mock function with given fields:
func (_m *MockClientAPI) GetBooks() ([]bookrepository.Book, error) {
	ret := _m.Called()

	var r0 []bookrepository.Book
	if rf, ok := ret.Get(0).(func() []bookrepository.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bookrepository.Book)
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

// MockClientAPI_GetBooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBooks'
type MockClientAPI_GetBooks_Call struct {
	*mock.Call
}

// GetBooks is a helper method to define mock.On call
func (_e *MockClientAPI_Expecter) GetBooks() *MockClientAPI_GetBooks_Call {
	return &MockClientAPI_GetBooks_Call{Call: _e.mock.On("GetBooks")}
}

func (_c *MockClientAPI_GetBooks_Call) Run(run func()) *MockClientAPI_GetBooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockClientAPI_GetBooks_Call) Return(_a0 []bookrepository.Book, _a1 error) *MockClientAPI_GetBooks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ReplaceUrl provides a mock function with given fields: url
func (_m *MockClientAPI) ReplaceUrl(url string) {
	_m.Called(url)
}

// MockClientAPI_ReplaceUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReplaceUrl'
type MockClientAPI_ReplaceUrl_Call struct {
	*mock.Call
}

// ReplaceUrl is a helper method to define mock.On call
//   - url string
func (_e *MockClientAPI_Expecter) ReplaceUrl(url interface{}) *MockClientAPI_ReplaceUrl_Call {
	return &MockClientAPI_ReplaceUrl_Call{Call: _e.mock.On("ReplaceUrl", url)}
}

func (_c *MockClientAPI_ReplaceUrl_Call) Run(run func(url string)) *MockClientAPI_ReplaceUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockClientAPI_ReplaceUrl_Call) Return() *MockClientAPI_ReplaceUrl_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockClientAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockClientAPI creates a new instance of MockClientAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockClientAPI(t mockConstructorTestingTNewMockClientAPI) *MockClientAPI {
	mock := &MockClientAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
