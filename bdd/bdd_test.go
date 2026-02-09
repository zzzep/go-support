package bdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewBDD(t *testing.T) {
	bdd := NewBDD(t)
	assert.NotNil(t, bdd)
	assert.Equal(t, t, bdd.T)
	assert.Empty(t, bdd.Scenarios)
}

func TestBDD_AddScenario(t *testing.T) {
	bdd := NewBDD(t)
	scenario := NewScenario()
	
	bdd.AddScenario(*scenario)
	assert.Len(t, bdd.Scenarios, 1)
	
	// Test chaining
	bdd.AddScenario(*scenario).AddScenario(*scenario)
	assert.Len(t, bdd.Scenarios, 3)
}

func TestBDD_Run_WithMultipleScenarios(t *testing.T) {
	bdd := NewBDD(t)
	
	scenario1 := NewScenario()
	scenario1.AddStep(func(sc Scenario, t *testing.T) {
		assert.True(t, true)
	})
	
	scenario2 := NewScenario()
	scenario2.AddStep(func(sc Scenario, t *testing.T) {
		assert.True(t, true)
	})
	
	bdd.AddScenario(*scenario1).AddScenario(*scenario2)
	bdd.Run()
}

func TestBDD_Run_ScenarioWithoutT(t *testing.T) {
	bdd := NewBDD(t)
	
	scenario := NewScenario()
	// Don't set scenario.T
	scenario.AddStep(func(sc Scenario, t *testing.T) {
		assert.True(t, true)
	})
	
	bdd.AddScenario(*scenario)
	bdd.Run()
}

func TestNewScenario(t *testing.T) {
	scenario := NewScenario()
	assert.NotNil(t, scenario)
	assert.Empty(t, scenario.Steps)
	assert.NotNil(t, scenario.Mocks)
	assert.Empty(t, scenario.Mocks)
}

func TestScenario_AddStep(t *testing.T) {
	scenario := NewScenario()
	
	step1 := func(sc Scenario, t *testing.T) {}
	step2 := func(sc Scenario, t *testing.T) {}
	
	scenario.AddStep(step1)
	assert.Len(t, scenario.Steps, 1)
	
	// Test chaining
	scenario.AddStep(step2).AddStep(step1)
	assert.Len(t, scenario.Steps, 3)
}

func TestScenario_AddMock(t *testing.T) {
	scenario := NewScenario()
	
	mock1 := &mock.Mock{}
	mock2 := &mock.Mock{}
	
	scenario.AddMock("key1", mock1)
	assert.Len(t, scenario.Mocks, 1)
	assert.Equal(t, mock1, scenario.Mocks["key1"])
	
	// Test chaining
	scenario.AddMock("key2", mock2).AddMock("key3", mock1)
	assert.Len(t, scenario.Mocks, 3)
}

func TestScenario_Run(t *testing.T) {
	scenario := NewScenario()
	scenario.T = t
	
	executed := false
	scenario.AddStep(func(sc Scenario, t *testing.T) {
		executed = true
		assert.True(t, true)
	})
	
	scenario.Run()
	assert.True(t, executed)
}

func TestScenario_Run_WithMultipleSteps(t *testing.T) {
	scenario := NewScenario()
	scenario.T = t
	
	step1Executed := false
	step2Executed := false
	
	scenario.AddStep(func(sc Scenario, t *testing.T) {
		step1Executed = true
	})
	
	scenario.AddStep(func(sc Scenario, t *testing.T) {
		step2Executed = true
	})
	
	scenario.Run()
	assert.True(t, step1Executed)
	assert.True(t, step2Executed)
}

// MockInterface for GenerateMock tests
type MockInterface interface {
	Method()
}

// MockWithInterface implements MockInterface
type MockWithInterface struct {
	mock.Mock
}

func (m *MockWithInterface) Method() {
	m.Called()
}

// MockWithoutInterface does NOT implement MockInterface
type MockWithoutInterface struct {
	mock.Mock
}

func TestGenerateMock_Success(t *testing.T) {
	mockInstance := &MockWithInterface{}
	
	// This should not panic
	_, m := GenerateMock[MockInterface](mockInstance)
	assert.NotNil(t, m)
	assert.Equal(t, mockInstance, m)
}

func TestGenerateMock_PanicWhenNotImplements(t *testing.T) {
	// MockWithoutInterface doesn't implement MockInterface
	mockInstance := &MockWithoutInterface{}
	
	// This should panic
	assert.Panics(t, func() {
		GenerateMock[MockInterface](mockInstance)
	}, "Mock does not implement interface")
}

