package bdd

import (
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/mock"
)

type Mock interface {
	On(method string, arguments ...interface{}) *mock.Call
	String() string
	TestData() objx.Map
	Test(t mock.TestingT)
	Called(arguments ...interface{}) mock.Arguments
	MethodCalled(methodName string, arguments ...interface{}) mock.Arguments
	AssertExpectations(t mock.TestingT) bool
	AssertNumberOfCalls(t mock.TestingT, methodName string, expectedCalls int) bool
	AssertCalled(t mock.TestingT, methodName string, arguments ...interface{}) bool
	AssertNotCalled(t mock.TestingT, methodName string, arguments ...interface{}) bool
	IsMethodCallable(t mock.TestingT, methodName string, arguments ...interface{}) bool
}
