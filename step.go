package main

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type BDDStep func(t *testing.T)

func (s *BDDStep) GetFunctionName(i interface{}) string {
	p := reflect.ValueOf(i).Pointer()
	n := runtime.FuncForPC(p).Name()
	if n == "" {
		panic("Function name is empty")
	}
	sp := strings.Split(n, ".")
	if len(sp) == 0 {
		panic("Function name does not contain a dot")
	}
	fn := sp[len(sp)-1]
	if len(fn) <= 2 {
		panic("Function invalid, declare it first")
	}
	return fn
}
