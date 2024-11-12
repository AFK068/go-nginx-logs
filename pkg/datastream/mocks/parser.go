// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Parser is an autogenerated mock type for the Parser type
type Parser[T any] struct {
	mock.Mock
}

type Parser_Expecter[T any] struct {
	mock *mock.Mock
}

func (_m *Parser[T]) EXPECT() *Parser_Expecter[T] {
	return &Parser_Expecter[T]{mock: &_m.Mock}
}

// Parse provides a mock function with given fields: line
func (_m *Parser[T]) Parse(line string) (*T, error) {
	ret := _m.Called(line)

	if len(ret) == 0 {
		panic("no return value specified for Parse")
	}

	var r0 *T
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*T, error)); ok {
		return rf(line)
	}
	if rf, ok := ret.Get(0).(func(string) *T); ok {
		r0 = rf(line)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*T)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(line)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Parser_Parse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Parse'
type Parser_Parse_Call[T any] struct {
	*mock.Call
}

// Parse is a helper method to define mock.On call
//   - line string
func (_e *Parser_Expecter[T]) Parse(line interface{}) *Parser_Parse_Call[T] {
	return &Parser_Parse_Call[T]{Call: _e.mock.On("Parse", line)}
}

func (_c *Parser_Parse_Call[T]) Run(run func(line string)) *Parser_Parse_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Parser_Parse_Call[T]) Return(_a0 *T, _a1 error) *Parser_Parse_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Parser_Parse_Call[T]) RunAndReturn(run func(string) (*T, error)) *Parser_Parse_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewParser creates a new instance of Parser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParser[T any](t interface {
	mock.TestingT
	Cleanup(func())
}) *Parser[T] {
	mock := &Parser[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
