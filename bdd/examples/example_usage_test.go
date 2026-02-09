package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zzzep/go-support/bdd"
)

// TestExampleBDDUsage demonstrates the complete usage of Given, When, Then types
// Demonstrates how to use scenario.Given(), scenario.When() and scenario.Then()
// without needing to write "AND" explicitly in functions
func TestExampleBDDUsage(t *testing.T) {
	scenario := bdd.NewScenario()

	// Add multiple Given steps - automatically treated as "AND"
	scenario.Given(ExIHaveAUser)
	scenario.Given(ExIHaveAProduct)    // Automatically "AND"
	scenario.Given(ExIAmAuthenticated) // Automatically "AND"

	// Add multiple When steps - automatically treated as "AND"
	scenario.When(ExIAddToCart)
	scenario.When(ExICompletePurchase) // Automatically "AND"

	// Add multiple Then steps - automatically treated as "AND"
	scenario.Then(ExCartHasProduct)
	scenario.Then(ExPurchaseWasCompleted) // Automatically "AND"

	bdd.NewBDD(t).
		AddScenario(*scenario).
		Run()
}

// TestExampleSimpleUsage demonstrates simple usage
func TestExampleSimpleUsage(t *testing.T) {
	scenario := bdd.NewScenario()

	// Simplified usage - just pass the method directly
	scenario.Given(ExIHaveAUser)
	scenario.When(ExIAddToCart)
	scenario.Then(ExCartHasProduct)

	bdd.NewBDD(t).
		AddScenario(*scenario).
		Run()
}

// TestExampleMixedUsage demonstrates usage mixing with traditional AddStep (compatibility)
func TestExampleMixedUsage(t *testing.T) {
	scenario := bdd.NewScenario()

	// Can mix the new methods with traditional AddStep
	scenario.Given(ExIHaveAUser)
	scenario.AddStep(ExIAddToCart)
	scenario.Then(ExCartHasProduct)

	bdd.NewBDD(t).
		AddScenario(*scenario).
		Run()
}

// Example steps - note that they don't need prefixes "Que", "Quando", "Então"
// The framework handles it automatically when there are multiple of the same type
// Logging is automatic, no need to call t.Log() manually

func ExIHaveAUser(sc bdd.Scenario, t *testing.T) {
	// Step implementation
}

func ExIHaveAProduct(sc bdd.Scenario, t *testing.T) {
	// Step implementation
}

func ExIAmAuthenticated(sc bdd.Scenario, t *testing.T) {
	// Step implementation
}

func ExIAddToCart(sc bdd.Scenario, t *testing.T) {
	// Step implementation
}

func ExICompletePurchase(sc bdd.Scenario, t *testing.T) {
	// Step implementation
}

func ExCartHasProduct(sc bdd.Scenario, t *testing.T) {
	assert.True(t, true)
}

func ExPurchaseWasCompleted(sc bdd.Scenario, t *testing.T) {
	assert.True(t, true)
}
