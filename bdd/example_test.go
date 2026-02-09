package bdd

import (
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeInterface interface{}
type fakeMock struct {
	mock.Mock
}

func (f *fakeMock) Method() {
	f.Called()
}

func TestPepe(t *testing.T) {
	scenario := NewScenario()
	f := &fakeMock{}
	scenario.AddMock("fakeMock", f)
	scenario.AddStep(DadoQueTenhaUmDado)
	scenario.AddStep(AndMockComSucesso)
	scenario.AddStep(QuandoVerificarSeTenhaOQuando)
	scenario.AddStep(EntaoDeveRetornarVerdadeiro)
	NewBDD(t).
		AddScenario(*scenario).
		Run()
}

func TestError(t *testing.T) {
	assert.Panics(t, func() {
		bdd := NewBDD(t)
		scenario := NewScenario()
		scenario.AddStep(func(_ Scenario, t *testing.T) {
			assert.True(t, false)
		})
		bdd.AddScenario(*scenario)
		bdd.Run()
	})
}

func TestExampleScenario(t *testing.T) {
	scenario := NewScenario()
	scenario.AddStep(Step1)
	scenario.AddStep(Step2)
	scenario.AddStep(Step3)
	NewBDD(t).
		AddScenario(*scenario).
		Run()
}

func Step1(sc Scenario, t *testing.T) {
	println("Example Step 1")
	assert.True(t, true)
}

func Step2(sc Scenario, t *testing.T) {
	println("Example Step 2")
	assert.True(t, true)
}

func Step3(sc Scenario, t *testing.T) {
	println("Example Step 3")
	assert.True(t, true)
}

func DadoQueTenhaUmDado(sc Scenario, t *testing.T) {
	println("Step 1")
}

func AndMockComSucesso(sc Scenario, t *testing.T) {
	sc.Mocks["fakeMock"].(*fakeMock).On("Method").Return().Once()
	sc.Mocks["fakeMock"].(*fakeMock).Method()
	println("Step Mock")
}

func QuandoVerificarSeTenhaOQuando(sc Scenario, t *testing.T) {
	println("Step 2")
}

func EntaoDeveRetornarVerdadeiro(sc Scenario, t *testing.T) {
	println("Step 3")
	assert.True(t, true)
}
