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
