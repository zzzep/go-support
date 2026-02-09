package bdd

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFunctionName_ValidFunction(t *testing.T) {
	var step Step = ValidTestFunction
	name := step.GetFunctionName(step)
	assert.Equal(t, "ValidTestFunction", name)
}

func ValidTestFunction(sc Scenario, t *testing.T) {
	// Valid function for testing
}

func TestGetFunctionName_AnotherValidFunction(t *testing.T) {
	var step Step = AnotherValidFunction
	name := step.GetFunctionName(step)
	assert.Equal(t, "AnotherValidFunction", name)
}

func AnotherValidFunction(sc Scenario, t *testing.T) {
	// Another valid function
}

func TestGetFunctionName_WithAnonymousFunction(t *testing.T) {
	anonFunc := func(sc Scenario, t *testing.T) {}
	var step Step = anonFunc
	
	// Anonymous functions should still work
	name := step.GetFunctionName(step)
	assert.NotEmpty(t, name)
	// Anonymous functions may have generated names, so we just check it's not empty
}

func TestGetFunctionName_WithMethodValue(t *testing.T) {
	receiver := &testReceiver{}
	var step Step = receiver.Method
	name := step.GetFunctionName(step)
	assert.NotEmpty(t, name)
}

type testReceiver struct{}

func (tr *testReceiver) Method(sc Scenario, t *testing.T) {
	// Method for testing
}

func TestGetFunctionName_EdgeCases(t *testing.T) {
	t.Run("Valid function with long name", func(t *testing.T) {
		var step Step = VeryLongFunctionNameForTestingPurposes
		name := step.GetFunctionName(step)
		assert.Equal(t, "VeryLongFunctionNameForTestingPurposes", name)
	})
}

func VeryLongFunctionNameForTestingPurposes(sc Scenario, t *testing.T) {
	// Long function name for testing
}

// Note: Testing panic cases for GetFunctionName is difficult because:
// 1. Empty function name - requires invalid function pointer (hard to create)
// 2. Function name without dot - requires special runtime conditions
// 3. Function name too short - requires function with <=2 character name (invalid Go syntax)
// These edge cases are protected by the function but are hard to test directly
// without manipulating runtime internals, which is not recommended.

func TestGetFunctionName_ReflectionEdgeCases(t *testing.T) {
	// Test that GetFunctionName works with different function types
	t.Run("Direct function assignment", func(t *testing.T) {
		var step Step = DirectFunction
		name := step.GetFunctionName(step)
		assert.Equal(t, "DirectFunction", name)
	})
	
	t.Run("Function pointer", func(t *testing.T) {
		var step Step = DirectFunction
		// Get function pointer
		funcPtr := reflect.ValueOf(step).Pointer()
		assert.NotZero(t, funcPtr)
	})
}

func DirectFunction(sc Scenario, t *testing.T) {
	// Direct function for testing
}
