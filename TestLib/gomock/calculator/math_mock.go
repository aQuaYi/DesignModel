// Code generated by MockGen. DO NOT EDIT.
// Source: math.go

// Package calculator is a generated GoMock package.
package calculator

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMath is a mock of Math interface
type MockMath struct {
	ctrl     *gomock.Controller
	recorder *MockMathMockRecorder
}

// MockMathMockRecorder is the mock recorder for MockMath
type MockMathMockRecorder struct {
	mock *MockMath
}

// NewMockMath creates a new mock instance
func NewMockMath(ctrl *gomock.Controller) *MockMath {
	mock := &MockMath{ctrl: ctrl}
	mock.recorder = &MockMathMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMath) EXPECT() *MockMathMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockMath) Add(arg0, arg1 int) int {
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockMathMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMath)(nil).Add), arg0, arg1)
}

// Sub mocks base method
func (m *MockMath) Sub(arg0 []int, arg1 *int) int {
	ret := m.ctrl.Call(m, "Sub", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// Sub indicates an expected call of Sub
func (mr *MockMathMockRecorder) Sub(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sub", reflect.TypeOf((*MockMath)(nil).Sub), arg0, arg1)
}
