package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPepe(t *testing.T) {
	scenario := NewScenario()
	scenario.AddStep(DadoQueTenhaUmDado)
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
		scenario.AddStep(func(t *testing.T) {
			assert.True(t, false)
		})
		bdd.AddScenario(*scenario)
		bdd.Run()
	})
}

func DadoQueTenhaUmDado(t *testing.T) {
	println("Step 1")
}

func QuandoVerificarSeTenhaOQuando(t *testing.T) {
	println("Step 2")
}

func EntaoDeveRetornarVerdadeiro(t *testing.T) {
	println("Step 3")
	assert.True(t, true)
}
