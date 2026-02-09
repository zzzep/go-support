package bdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenWhenThen(t *testing.T) {
	scenario := NewScenario()

	// Using the new types - no need to write "AND" explicitly
	scenario.Given(GivenStep1)
	scenario.Given(GivenStep2) // Automatically treated as "AND"
	scenario.When(WhenStep1)
	scenario.When(WhenStep2) // Automatically treated as "AND"
	scenario.Then(ThenStep1)
	scenario.Then(ThenStep2) // Automatically treated as "AND"

	NewBDD(t).
		AddScenario(*scenario).
		Run()
}

func TestMixedSteps(t *testing.T) {
	scenario := NewScenario()

	// Can mix the new types with traditional AddStep
	scenario.Given(GivenStep1)
	scenario.AddStep(Step(WhenStep1))
	scenario.Then(ThenStep1)

	NewBDD(t).
		AddScenario(*scenario).
		Run()
}

// Example steps
func GivenStep1(sc Scenario, t *testing.T) {
	t.Log("Given: Step 1")
}

func GivenStep2(sc Scenario, t *testing.T) {
	t.Log("And: Step 2 (automatic)")
}

func WhenStep1(sc Scenario, t *testing.T) {
	t.Log("When: Step 1")
}

func WhenStep2(sc Scenario, t *testing.T) {
	t.Log("And: Step 2 (automatic)")
}

func ThenStep1(sc Scenario, t *testing.T) {
	t.Log("Then: Step 1")
	assert.True(t, true)
}

func ThenStep2(sc Scenario, t *testing.T) {
	t.Log("And: Step 2 (automatic)")
	assert.True(t, true)
}
