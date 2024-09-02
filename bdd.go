package main

import "testing"

type BDD struct {
	T         *testing.T
	Scenarios []BDDScenario
}

func NewBDD(t *testing.T) *BDD {
	return &BDD{
		T: t,
	}
}

func (b *BDD) AddScenario(scenario BDDScenario) *BDD {
	b.Scenarios = append(b.Scenarios, scenario)
	return b
}

func (b *BDD) Run() {
	for _, scenario := range b.Scenarios {
		if scenario.T == nil {
			scenario.T = b.T
		}
		scenario.Run()
	}
}
