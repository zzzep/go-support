package bdd

import "testing"

// Given represents a Given step in BDD
type Given Step

// When represents a When step in BDD
type When Step

// Then represents a Then step in BDD
type Then Step

// ToStep converts Given to Step
func (g Given) ToStep() Step {
	return Step(g)
}

// ToStep converts When to Step
func (w When) ToStep() Step {
	return Step(w)
}

// ToStep converts Then to Step
func (t Then) ToStep() Step {
	return Step(t)
}

// GetFunctionName returns the function name for Given
func (g Given) GetFunctionName(i interface{}) string {
	return GetFunctionName(Step(g))
}

// GetFunctionName returns the function name for When
func (w When) GetFunctionName(i interface{}) string {
	return GetFunctionName(Step(w))
}

// GetFunctionName returns the function name for Then
func (tn Then) GetFunctionName(i interface{}) string {
	return GetFunctionName(Step(tn))
}

// Call executes the Given step
func (g Given) Call(sc Scenario, t *testing.T) {
	Step(g)(sc, t)
}

// Call executes the When step
func (w When) Call(sc Scenario, t *testing.T) {
	Step(w)(sc, t)
}

// Call executes the Then step
func (tn Then) Call(sc Scenario, t *testing.T) {
	Step(tn)(sc, t)
}

// NewGiven converts a Step to Given type
func NewGiven(step Step) Given {
	return Given(step)
}

// NewWhen converts a Step to When type
func NewWhen(step Step) When {
	return When(step)
}

// NewThen converts a Step to Then type
func NewThen(step Step) Then {
	return Then(step)
}
