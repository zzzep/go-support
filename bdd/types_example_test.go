package bdd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zzzep/go-support/bdd"
)

// TestExampleGivenWhenThen demonstrates the usage of the new Given, When, Then types
// No need to write "And..." in functions anymore - the framework handles it automatically
func TestExampleGivenWhenThen(t *testing.T) {
	scenario := bdd.NewScenario()

	// Add multiple Given steps - automatically treated as "AND"
	scenario.Given(IHaveAUser)
	scenario.Given(IHaveAProduct)    // Automatically "AND"
	scenario.Given(IAmAuthenticated) // Automatically "AND"

	// Add multiple When steps - automatically treated as "AND"
	scenario.When(IAddToCart)
	scenario.When(ICompletePurchase) // Automatically "AND"

	// Add multiple Then steps - automatically treated as "AND"
	scenario.Then(CartHasProduct)
	scenario.Then(PurchaseWasCompleted) // Automatically "AND"

	bdd.NewBDD(t).
		AddScenario(*scenario).
		Run()
}

// TestExampleUsage demonstrates simple usage with helpers
func TestExampleUsage(t *testing.T) {
	scenario := bdd.NewScenario()

	// Simplified usage - just pass the method directly
	scenario.Given(IHaveAUser)
	scenario.When(IAddToCart)
	scenario.Then(CartHasProduct)

	bdd.NewBDD(t).
		AddScenario(*scenario).
		Run()
}

// Example steps - note that they don't need prefixes "Que", "Quando", "Então"
// Logging is automatic, no need to call t.Log() manually
func IHaveAUser(sc bdd.Scenario, t *testing.T) {
	// Step implementation
}

func IHaveAProduct(sc bdd.Scenario, t *testing.T) {
	// Step implementation
}

func IAmAuthenticated(sc bdd.Scenario, t *testing.T) {
	// Step implementation
}

func IAddToCart(sc bdd.Scenario, t *testing.T) {
	// Step implementation
}

func ICompletePurchase(sc bdd.Scenario, t *testing.T) {
	// Step implementation
}

func CartHasProduct(sc bdd.Scenario, t *testing.T) {
	assert.True(t, true)
}

func PurchaseWasCompleted(sc bdd.Scenario, t *testing.T) {
	assert.True(t, true)
}
