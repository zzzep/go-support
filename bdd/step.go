package bdd

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type Step func(sc Scenario, t *testing.T)

func GetFunctionName(step Step) string {
	p := reflect.ValueOf(step).Pointer()
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

// GetFunctionName is kept for backward compatibility
func (s *Step) GetFunctionName(i interface{}) string {
	if step, ok := i.(Step); ok {
		return GetFunctionName(step)
	}
	panic("Invalid step type")
}
